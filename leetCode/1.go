package main

import "fmt"

/**
1. Two Sum
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.
*/
func main() {
	nums := []int{3, 4, 7, 9}
	target := 16
	twoSum := twoSum(nums, target)
	fmt.Println(twoSum)
}

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)

	for key, num := range nums {
		for i := key + 1; i < lenNums; i++ {
			if i == lenNums {
				break
			}
			if num+nums[i] == target {
				return []int{key, i}
			}
		}

	}

	return []int{0, 0}
}
