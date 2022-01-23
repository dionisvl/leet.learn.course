package main

import (
	"fmt"
)

/**
10 Даны три натуральных числа a, b, c.
Определите, существует ли треугольник с такими сторонами. Sample Input: 3 3 3 Sample Output: Да
*/
func main() {
	var a = 5
	var b = 1
	var c = 5
	answer := isTriangleOrNot(a, b, c)
	fmt.Println(answer)
}

func isTriangleOrNot(a int, b int, c int) string {
	if a+b > c {
		return "Да"
	}

	return "Нет"
}
