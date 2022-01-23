package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
14 Калькулятор: Дано действие и числа с которыми эти действия нужно совершить. Выведите результат.
Входные данные: Действие(строка). Например: "сложение", "корень", "степень" Натуральное число 1 Натуральное число 2
Выходные данные: Результат операции
Sample Input: "корень" 4 Sample Output: 2
Sample Input: "сложение" 2 2 Sample Output: 4
*/
func main() {
	var action string
	var num1 int
	var num2 int

	fmt.Print("Enter action: ")
	scanLn, err := fmt.Scanln(&action)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	if !isCorrectAction(action) {
		fmt.Println("не верно введено действие!")
		return
	}

	fmt.Print("Enter num 1: ")
	scanLn, err = fmt.Scanln(&num1)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	if strings.TrimRight(action, "\n") == "корень" || action == "степень" {
		num2 = 0
	} else {
		fmt.Print("Enter num 2: ")
		scanLn, err = fmt.Scanln(&num2)
		if err != nil {
			fmt.Print(scanLn, err)
		}
	}

	answer := calc(action, num1, num2)
	fmt.Println("result: ", answer)
}

func calc(action string, num1 int, num2 int) string {
	var answer2 string
	switch action {
	case "сложение":
		answer2 = strconv.Itoa(num1 + num2)
	case "вычитание":
		answer2 = strconv.Itoa(num1 - num2)
	case "корень":
		answer2 = strconv.FormatFloat(math.Sqrt(float64(num1)), 'E', -1, 64)
	case "степень":
		answer2 = strconv.Itoa(num1 * num1)
	default:
		answer2 = "не верно введено действие!"
	}

	return answer2
}

func isCorrectAction(inputStr string) bool {
	goodActions := []string{"сложение", "вычитание", "корень", "степень"}
	for _, v := range goodActions {
		if v == inputStr {
			return true
		}
	}

	return false
}
