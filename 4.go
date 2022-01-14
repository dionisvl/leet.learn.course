package main

import (
	"fmt"
	"strconv"
)

func main() {
	first := getFirstNumber(521)
	fmt.Println(first)
}

func getFirstNumber(num int) int {
	str := strconv.Itoa(num)

	firstNum, err := strconv.Atoi(string(str[0]))

	if err != nil {
		fmt.Println(err)
	}

	return firstNum
}
