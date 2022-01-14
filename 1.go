package main

import (
	"fmt"
	"strconv"
)

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
