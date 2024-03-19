/*
Задача № 4. Проверить, является ли четырехзначное число палиндромом
Пример:
Вход: 1221  Выход: 1221 - палиндром
Вход: 1234  Выход: 1234 - не палиндром
*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter the number: ")
	fmt.Scanf("%d%d\n", &num)

	var a, b, c, d int

	a = num / 1000
	b = (num % 1000) / 100
	c = (num % 100) / 10
	d = (num % 10) % 10

	if a == d && b == c {
		fmt.Printf("Число %d палиндром\n", num)
	} else {
		fmt.Printf("Число %d не палиндром\n ", num)
	}
}
