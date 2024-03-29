/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Order represents a customer's order.
type Order struct {
	Products []*Product
	*Customer
	*Address
}

// NewOrder creates a new Order.
func NewOrder() *Order {
	return &Order{}
}

// Print prints the order details to the console.
func (o *Order) Print() {
	fmt.Printf("\n\nДанные заказа\n\n")
	fmt.Println("ФИО покупателя:", o.Customer.Name)
	fmt.Println("Контактный телефон:", o.Customer.Phone)
	fmt.Println("Адрес:", o.Address.ZipCode, o.Address.City, o.Address.Street, o.Address.House, o.Address.Apartment)
	for index, p := range o.Products {
		fmt.Printf("Товар № %d:\n", index+1)
		fmt.Println("\tНаименование товара:", p.Name)
		fmt.Println("\tКоличество:", p.Quantity)
	}
}

// AddProduct adds a product to the order.
func (o *Order) AddProduct(p *Product) {
	o.Products = append(o.Products, p)
}

// Product represents a product in an order.
type Product struct {
	Name     string
	Quantity int
}

// NewProduct creates a new Product.
func NewProduct() *Product {
	return &Product{}
}

// Customer represents a customer placing an order.
type Customer struct {
	Name  string
	Phone string
}

// NewCustomer creates a new Customer.
func NewCustomer() *Customer {
	return &Customer{}
}

// Address represents a customer's shipping address.
type Address struct {
	ZipCode   int
	City      string
	Street    string
	House     string
	Apartment int
}

// NewAddress creates a new Address.
func NewAddress() *Address {
	return &Address{}
}

func main() {
	order := NewOrder()
	customer := NewCustomer()
	address := NewAddress()

	order.Customer = customer
	order.Address = address

	readCustomer(customer)
	readAddress(address)
	readProducts(order)

	order.Print()
}

func readCustomer(customer *Customer) {
	for {
		text, err := readString("ФИО покупателя: ", checkCustomerName)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		customer.Name = text
		break
	}

	for {
		text, err := readString("Телефон покупателя: ", checkCustomerPhone)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		customer.Phone = text
		break
	}
}

func readAddress(address *Address) {
	for {
		num, err := readInt("Адрес покупателя. Индекс: ", checkZipCode)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		address.ZipCode = num
		break
	}

	text, _ := readString("Город: ", nil)
	address.City = text

	text, _ = readString("Улица: ", nil)
	address.Street = text

	text, _ = readString("Дом: ", nil)
	address.House = text

	for {
		num, err := readInt("Квартира: ", nil)
		if err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			continue
		}
		address.Apartment = num
		break
	}
}

func readProducts(order *Order) {
	for i := 1; i > 0; i++ {
		product := NewProduct()
		fmt.Printf("Заполнение товара %d\n", i)

		for {
			text, err := readString("Наименование: ", checkProductName)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				continue
			}
			product.Name = text
			break
		}

		for {
			num, err := readInt("Количество: ", checkProductQuantity)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				continue
			}
			product.Quantity = num
			break
		}

		order.AddProduct(product)

		text, _ := readString("Добавить еще товар? (y/N): ", nil)
		if text != "y" {
			break
		}
	}
}

func checkZipCode(zipCode int) error {
	if utf8.RuneCountInString(fmt.Sprint(zipCode)) != 6 {
		return fmt.Errorf("индекс должен состоять из %d цифр", 6)
	}
	return nil
}

func checkCustomerPhone(phone string) error {
	if regexp.MustCompile(`[a-z]`).MatchString(phone) {
		return fmt.Errorf("телефон должен содержать только цифры")
	}
	if utf8.RuneCountInString(phone) != 10 {
		return fmt.Errorf("телефон должен состоять из %d цифр", 10)
	}

	return nil
}

func checkCustomerName(name string) error {
	if regexp.MustCompile(`\d`).MatchString(name) {
		return fmt.Errorf("нельзя использовать цифры в ФИО")
	}
	return nil
}

func checkProductQuantity(quantity int) error {
	if quantity <= 0 {
		return fmt.Errorf("количество товара должно быть больше %d", 0)
	}
	return nil
}

func checkProductName(name string) error {
	if utf8.RuneCountInString(name) < 1 || utf8.RuneCountInString(name) > 100 {
		return fmt.Errorf("наименование товара должно быть от %d до %d символов", 1, 100)
	}
	return nil
}

func readInt(text string, validator func(int) error) (i int, err error) {
	fmt.Print(text)
	str := readAnswer()

	i, err = strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("значение должно состоять только из цифр")
	}

	if validator != nil {
		err = validator(i)
		if err != nil {
			return 0, err
		}
	}

	return i, nil
}

func readString(text string, validator func(string) error) (str string, err error) {
	fmt.Print(text)
	str = readAnswer()

	if validator != nil {
		err = validator(str)
		if err != nil {
			return "", err
		}
	}

	return
}

func readAnswer() string {
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		text := strings.TrimSpace(input.Text())
		return text
	}
	return ""
}
