package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := sum(36)
	fmt.Println(sum)
}

/* sum of all numbers between first number and second */
func sum(twoNumbers uint8) uint8 {
	str := strconv.Itoa(int(twoNumbers))
	var sum uint8 = 0

	firstNum, err := strconv.Atoi(string(str[0]))
	secondNum, err := strconv.Atoi(string(str[1]))
	if err != nil {
		fmt.Println(err)
	}

	for i := firstNum + 1; i < secondNum; i++ {
		sum = sum + uint8(i)
	}

	return sum
}
