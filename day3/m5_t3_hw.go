/*
Задача №2
Вход:
Пользователь должен ввести правильный пароль, состоящий из:
цифр,
букв латинского алфавита(строчные и прописные) и
специальных символов  special = "_!@#$%^&"

Всего 4 набора различных символов.
В пароле обязательно должен быть хотя бы один символ из каждого набора.
Длина пароля от 8(мин) до 15(макс) символов.
Максимальное количество попыток ввода неправильного пароля - 5.
Каждый раз выводим номер попытки.
*Желательно выводить пояснение, почему пароль не принят и что нужно исправить.

digits = "0123456789"
lowercase = "abcdefghiklmnopqrstvxyz"
uppercase = "ABCDEFGHIKLMNOPQRSTVXYZ"
special = "_!@#$%^&"

Выход:
Написать, что ввели правильный пароль.

Пример:
хороший пароль -> o58anuahaunH!
хороший пароль -> aaaAAA111!!!
плохой пароль -> saucacAusacu8
*/

// Package main contains the entry point for the application.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	// attemptLimit is the maximum number of attempts allowed to enter a valid password.
	attemptLimit = 5
	// minLength is the minimum length of a valid password.
	minLength = 8
	// maxLength is the maximum length of a valid password.
	maxLength = 15
	// digits is a string containing all digits (0-9).
	digits = "0123456789"
	// lowercase is a string containing all lowercase letters (a-z).
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	// uppercase is a string containing all uppercase letters (A-Z).
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// special is a string containing all special characters (!@#$%^&).
	special = "_!@#$%^&"
)

// main is the entry point for the application.
func main() {
	fmt.Println("Ввод пароля")
	for i := 1; i <= attemptLimit; i++ {
		input := bufio.NewScanner(os.Stdin)
		fmt.Printf("Попытка ввода №%d: ", i)
		if input.Scan() {
			password := strings.TrimSpace(input.Text())
			ok, errorsSlice := checkPassword(password)
			if ok {
				fmt.Printf("Вы ввели правильный пароль c %d попытки!\n", i)
				break
			}

			if i == attemptLimit {
				fmt.Println("Слишком много попыток ввода неправильного пароля. Попробуйте позже.")
				return
			}

			fmt.Printf("Исправьте ошибки:\n")
			for _, val := range errorsSlice {
				fmt.Printf("- %s\n", val)
			}
		}
	}
}

// checkPassword checks if the given password is valid.
func checkPassword(password string) (bool, []error) {
	errorsSlice := make([]error, 0, 5)

	checkers := [5]func(string) error{checkLength, checkDigits, checkLowercase, checkUppercase, checkSpecialChars}

	for _, checker := range checkers {
		if err := checker(password); err != nil {
			errorsSlice = append(errorsSlice, err)
		}
	}

	return len(errorsSlice) == 0, errorsSlice
}

// checkLength checks if the given password is of the correct length.
func checkLength(password string) error {
	if utf8.RuneCountInString(password) < minLength || utf8.RuneCountInString(password) > maxLength {
		return fmt.Errorf("длина пароля должна быть от %d до %d символов", minLength, maxLength)
	}
	return nil
}

// checkDigits checks if the given password contains at least one digit.
func checkDigits(password string) error {
	if !strings.ContainsAny(password, digits) {
		return errors.New("пароль должен содержать хотя бы одну цифру (0-9)")
	}
	return nil
}

// checkLowercase checks if the given password contains at least one lowercase letter.
func checkLowercase(password string) error {
	if !strings.ContainsAny(password, lowercase) {
		return errors.New("пароль должен содержать хотя бы одну букву нижнего регистра (a-z)")
	}
	return nil
}

// checkUppercase checks if the given password contains at least one uppercase letter.
func checkUppercase(password string) error {
	if !strings.ContainsAny(password, uppercase) {
		return errors.New("пароль должен содержать хотя бы одну букву верхнего регистра (A-Z)")
	}
	return nil
}

// checkSpecialChars checks if the given password contains at least one special character.
func checkSpecialChars(password string) error {
	if !strings.ContainsAny(password, special) {
		return fmt.Errorf("пароль должен содержать хотя бы один специальный символ (%s)", special)
	}
	return nil
}
