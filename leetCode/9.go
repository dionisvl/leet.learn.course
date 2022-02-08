package main

import (
	"fmt"
	"strconv"
)

/**
Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
For example, 121 is a palindrome while 123 is not.
*/
func main() {
	result := palindrome(1235321)
	fmt.Println(result)
}

func palindrome(x int) bool {
	var str = strconv.Itoa(x)
	var strLen = len(str)

	for i := 0; i <= strLen/2; i++ {
		if str[i] == str[strLen-1-i] {
			continue
		}

		return false
	}

	return true
}
