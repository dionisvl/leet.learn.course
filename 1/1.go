package main

import (
	"fmt"
	"strconv"
)

/**
1 Дано натуральное число, выведите его последнюю цифру.
Входные данные: На вход дается натуральное число N, не превосходящее 10000.
Выходные данные: целое число. Sample Input: 567 Sample Output: 7 Sample Input: 856 Sample Output: 6
*/
func main() {
	x := 669
	lastNum := getLastNum(x)
	fmt.Println(lastNum)
}

func getLastNum(num int) int {
	t := strconv.Itoa(num)

	lastNum, err := strconv.Atoi(t[len(t)-1:])
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	return lastNum
}
