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

	if a > b {
		changeVals(&a, &b)
	}
	if b > c {
		changeVals(&b, &c)
	}
	if a > b {
		changeVals(&a, &b)
	}
	fmt.Printf("Result: %d%d%d\n", a, b, c)
}

func changeVals(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
