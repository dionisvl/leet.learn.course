package main

import (
	"fmt"
)

/**
Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.
For example, 121 is a palindrome while 123 is not.
*/
func main() {
	result := isPalindrome(1331)
	fmt.Println(result)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var intSlice = IntToSlice(int64(x), []int64{})
	var strLen = len(intSlice)

	for i := 0; i < strLen/2; i++ {
		//fmt.Println(intSlice[i])
		//fmt.Println(intSlice[strLen-1-i])
		if intSlice[i] == intSlice[strLen-1-i] {
			continue
		}

		return false
	}

	return true
}

func IntToSlice(n int64, sequence []int64) []int64 {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int64{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}
