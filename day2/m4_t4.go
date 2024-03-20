/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	printChar = 0
)

func main() {
	boardSize := readBoardSize()
	printBoard(boardSize)
}

func readBoardSize() (boardSize int) {
	var err error
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите число от 0 до 20: ")
		if input.Scan() {
			inputString := input.Text()
			boardSize, err = strconv.Atoi(inputString)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				continue
			}
		}
		if boardSize >= 0 && boardSize <= 20 {
			break
		}
	}
	return
}

func printBoard(boardSize int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			fmt.Printf("%d ", printChar)
			printChar = printChar ^ 1
		}
		if boardSize%2 == 0 {
			printChar = printChar ^ 1
		}
		fmt.Print("\n")
	}
}
