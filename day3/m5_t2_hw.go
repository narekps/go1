/*
Задача №1
Написать функцию, которая расшифрует строку.
code = "220411112603141304"
Каждые две цифры - это либо буква латинского алфавита в нижнем регистре либо пробел.
Отчет с 00 -> 'a' и до 25 -> 'z', 26 -> ' '(пробел).
Вход: строка из цифр. Выход: Текст.
Проверка работы функции выполняется через вторую строку.

codeToString(code) -> "???????'
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var alphabet = map[string]string{
	"00": "a",
	"01": "b",
	"02": "c",
	"03": "d",
	"04": "e",
	"05": "f",
	"06": "g",
	"07": "h",
	"08": "i",
	"09": "j",
	"10": "k",
	"11": "l",
	"12": "m",
	"13": "n",
	"14": "o",
	"15": "p",
	"16": "q",
	"17": "r",
	"18": "s",
	"19": "t",
	"20": "u",
	"21": "v",
	"22": "w",
	"23": "x",
	"24": "y",
	"25": "z",
	"26": " ",
}

func main() {
	codeStr, err := readLine()
	if err != nil {
		fmt.Println(err)
		return
	}
	decodedString := codeToString(codeStr)
	fmt.Printf("Ответ: \"%s\"\n", decodedString)
}

// readLine reads a line of input from the user.
func readLine() (string, error) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите код для расшифровки: ")
	if input.Scan() {
		inputCode := strings.TrimSpace(input.Text())
		err := validateInput(inputCode)
		if err != nil {
			return "", err
		}
		return inputCode, nil
	}
	return "", errors.New("не удалось прочитать строку")
}

// validateInput checks if the input code is valid.
func validateInput(inputCode string) error {
	// Check if the input code is empty.
	if utf8.RuneCountInString(inputCode) == 0 {
		return errors.New("входная строка не должна быть пустой")
	}

	// Check if the input code has an odd number of characters.
	if utf8.RuneCountInString(inputCode)%2 != 0 {
		return errors.New("входная строка должна иметь четное количество символов")
	}

	// Check if the input code contains only digits.
	for i := 0; i < utf8.RuneCountInString(inputCode); i++ {
		_, err := strconv.Atoi(string(inputCode[i]))
		if err != nil {
			return fmt.Errorf("входная строка содержит запрещенный символ: \"%s\"", string(inputCode[i]))
		}

	}

	return nil
}

// codeToString decodes a string of digits into a human-readable string.
func codeToString(inputCode string) (result string) {
	for i := 0; i < utf8.RuneCountInString(inputCode); i += 2 {
		code := string(inputCode[i]) + string(inputCode[i+1])
		result += alphabet[code]
	}
	return
}
