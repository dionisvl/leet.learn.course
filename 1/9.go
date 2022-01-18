package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num1 = 35
	var num2 = 39568
	coincidences := twoNumbersCoincidences(num1, num2)
	fmt.Println(coincidences)
}

/**
9 Даны два числа. Определить цифры, входящие в запись обоих чисел.
Sample Input: 35 39568 Sample Output: 3 5
*/
func twoNumbersCoincidences(num1 int, num2 int) string {
	var str1 = strconv.Itoa(num1)
	var str2 = strconv.Itoa(num2)
	var answer = ""

	for i := 0; i < len(str1); i++ {
		if strings.Contains(str2, string(str1[i])) == true {
			answer += " " + string(str1[i])
		}
	}

	return answer
}
