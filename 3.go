package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	count := getNoRepeatableNumbersCount(121234555)
	fmt.Println(count)
}

func getNoRepeatableNumbersCount(num int) int {
	str := strconv.Itoa(num)
	count := 0 // кол-во не повторяющихся цифр

	for i := 0; i < len(str)-1; i++ {
		if strings.Count(str, string(str[i])) == 1 {
			count++
		}
	}

	return count
}
