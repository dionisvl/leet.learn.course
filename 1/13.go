package main

import (
	"fmt"
	"strconv"
)

/**
13 Счастливый билет: Определите является ли билет счастливым.
Входные данные: На вход программе подается целое число N
Выходные данные: Выведите "Да", если или "Нет".
Sample Input: 111003 Sample Output: Да
*/
func main() {
	var ticketNumber = 111003
	answer := isHappyTicket(ticketNumber)
	fmt.Println(answer)
}

func isHappyTicket(ticketNumber int) string {
	var str = strconv.Itoa(ticketNumber)
	var half = len(str) / 2
	var sumOfFirstPart = 0
	var sumOfLastPart = 0

	//calc first part of ticket:
	for i := 0; i <= len(str)-1; i++ {
		if i >= half {
			break
		}
		num, err := strconv.Atoi(string(str[len(str)-(len(str)-i)]))

		if err != nil {
			fmt.Println(err)
		}

		sumOfFirstPart += num
	}

	//calc last part of ticket:
	for i := 0; i <= len(str)-1; i++ {
		if i >= half {
			break
		}
		num, err := strconv.Atoi(string(str[len(str)-(i+1)]))

		if err != nil {
			fmt.Println(err)
		}

		sumOfLastPart += num
	}

	if sumOfLastPart == sumOfFirstPart {
		return "да"
	}

	return "нет"
}
