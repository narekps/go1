/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите количество бутылок от 0 до 200: ")
		if input.Scan() {
			inputString := input.Text()
			num, err := strconv.Atoi(inputString)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				continue
			}
			if num < 0 || num > 200 {
				fmt.Printf("Ошибка: %s\n", "Введите учисло от 0 до 200")
				continue
			}
			str := calcCount(num)
			fmt.Printf("%d %s\n", num, str)
		}
	}
}

func calcCount(num int) string {
	switch true {
	case num%100 >= 10 && num%100 <= 20:
		return "бутылок"
	case num%10 == 1:
		return "бутылка"
	case num%10 >= 2 && num%10 <= 4:
		return "бутылки"
	default:
		return "бутылок"
	}
}
