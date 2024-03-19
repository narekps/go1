/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*/

package main

import "fmt"

const (
	price = 48
)

func main() {
	var distance, consumption int
	var sum float64

	fmt.Print("Введите расстояние (50 - 10000 км): ")
	fmt.Scanf("%d\n", &distance)

	if distance < 50 || distance > 10000 {
		fmt.Println("Расстояние указано неправильно.")
		return
	}

	fmt.Print("Введите расход в литрах на 100 км (5 - 25 литров): ")
	fmt.Scanf("%d\n", &consumption)

	if consumption < 5 || consumption > 25 {
		fmt.Println("Расход топлива указано неправильно.")
		return
	}

	sum = float64(distance) / 100 * float64(consumption) * price

	fmt.Printf("Стоимость поездки: %.2f руб. при стоимости литра %d руб.\n", sum, price)
}
