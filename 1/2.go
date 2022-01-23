package main

import (
	"fmt"
	"strconv"
)

/**
2 Дано целое число. Найдите число десятков
Входные данные: На вход дается натуральное число N, не превосходящее 10000.
Выходные данные: целое число.
Sample Input: 567 Sample Output: 6
Sample Input: 856 Sample Output: 5
*/
func main() {
	countTens := getCountTens(31415966)
	fmt.Println(countTens)
}

func getCountTens(num int) int {
	str := strconv.Itoa(num)
	str2 := string(str[len(str)-2])

	countTens, err := strconv.Atoi(str2)
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	return countTens
}
