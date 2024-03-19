/*
Задача № 3. Вывести на экран в порядке возрастания три введенных числа
Пример:
Вход: 1, 9, 2
Выход: 1, 2, 9
*/

package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d\n", &a)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d\n", &b)
	fmt.Print("Enter the third number: ")
	fmt.Scanf("%d\n", &c)

	if a <= b && b <= c {
		fmt.Printf("Result: %d%d%d\n", a, b, c)
		return
	}

	if a <= c && c <= b {
		fmt.Printf("Result: %d%d%d\n", a, c, b)
		return
	}

	if b <= a && a <= c {
		fmt.Printf("Result: %d%d%d\n", b, a, c)
		return
	}

	if b <= c && c <= a {
		fmt.Printf("Result: %d%d%d\n", b, c, a)
		return
	}

	if c <= a && a <= b {
		fmt.Printf("Result: %d%d%d\n", c, a, b)
		return
	}

	if c <= b && b <= a {
		fmt.Printf("Result: %d%d%d\n", c, b, a)
		return
	}
}
