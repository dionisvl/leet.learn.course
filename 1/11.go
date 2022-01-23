package main

import (
	"fmt"
	"strconv"
)

/**
11 Найдите цифровой корень числа
Sample Input: 3333 Sample Output: 3
*/
func main() {
	var num = 1237
	answer := digitalRootOfNumber(num)
	fmt.Println(answer)
}

func digitalRootOfNumber(num int) int {
	var str = strconv.Itoa(num)
	var sumOfNumbers = 0
	var countNums = len(str)

	if countNums > 1 {
		for i := 0; i < len(str); i++ {
			num, err := strconv.Atoi(string(str[len(str)-(i+1)]))

			if err != nil {
				fmt.Println(err)
			}

			sumOfNumbers += num
		}
		if len(strconv.Itoa(sumOfNumbers)) > 1 {
			sumOfNumbers = digitalRootOfNumber(sumOfNumbers)
		}
	}

	return sumOfNumbers
}
