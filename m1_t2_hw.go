/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 100, выход: 001
*/

package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter the number: ")
	fmt.Scanf("%d\n", &num)

	var a, b, c int

	a = (num - (num % 100)) / 100
	b = (num % 100) / 10
	c = (num % 100) % 10

	fmt.Printf("Result: %d%d%d\n", c, b, a)
}
