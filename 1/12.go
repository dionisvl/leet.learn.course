package main

import (
	"fmt"
	"math"
)

/**
12 Дано натуральное число A > 1.
Определите, каким по счету числом Фибоначчи оно является.
Если оно не является числом Фибоначчи вывести 0.
Sample Input: 8
*/
func main() {
	var a = 55
	answer := fiboNum(a)
	fmt.Println(answer)
}

func fiboNum(num int) int {
	if !isFibonacci(num) {
		return 0
	}

	var prev = 0
	var curr = 1
	var fib = 0

	for i := 0; i <= 200; i++ {

		prev = curr
		curr = fib
		fib = prev + curr

		if num == curr {
			return i
		}
	}

	return 0
}

func isPerfectSquare(x int) bool {
	var s = math.Sqrt(float64(x))
	return s*s == float64(x)
}

func isFibonacci(n int) bool {
	return isPerfectSquare(5*n*n+4) || isPerfectSquare(5*n*n-4)
}
