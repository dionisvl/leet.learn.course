package main

import (
	"fmt"
	"strconv"
)

func main() {
	countTens := getNoRepeatableNumbersCount(31415926)
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
