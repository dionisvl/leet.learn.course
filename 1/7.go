package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := sumOfNumbersDivisibleByEight()
	fmt.Println(sum)
}

func sumOfNumbersDivisibleByEight() uint32 {
	var sum uint32 = 0

	fmt.Print("Enter count of numbers: ")
	var countNumbers int
	num, err := fmt.Scanln(&countNumbers)

	var iterableNumber uint32 = 0

	for i := 1; i <= countNumbers; i++ {
		fmt.Print("Enter number № " + strconv.Itoa(i) + ": ")

		num, err = fmt.Scanln(&iterableNumber)
		if err != nil {
			fmt.Print(num, err)
		}

		if iterableNumber%8 == 0 {
			sum = sum + iterableNumber
		}
	}

	return sum
}
