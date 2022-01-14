package main

import (
	"fmt"
)

func main() {
	leap := isYearLeap(3)
	fmt.Println(leap)
}

func isYearLeap(year int) string {

	answer := "net"

	if (year%4 == 0) && ((year%100 != 0) || (year%400 == 0)) {
		answer = "Да"
	}

	return answer
}
