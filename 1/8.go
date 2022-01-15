package main

import (
	"fmt"
)

func main() {
	sum := countOfMaxNumbers()
	fmt.Println(sum)
}

/**
8 Последовательность состоит из натуральных чисел и завершается числом 0.
Определите количество элементов этой последовательности,
которые равны ее наибольшему элементу. Sample Input: 5 6 9 8 9 7 0 Sample Output: 2
*/
func countOfMaxNumbers() int {
	var str []int                               // строка для наглядного вывода последовательности на экран
	var similarNumsCounts = make(map[uint]uint) // количество идентичных чисел
	var enteringNum uint = 0                    // число, которое будет вводиться с консоли в последовательность

	for i := 1; i <= 10000; i++ {
		fmt.Print("Enter number: ")

		num, err := fmt.Scanln(&enteringNum)
		if err != nil {
			fmt.Print(num, err)
		}

		if enteringNum == 0 {
			break
		}

		if similarNumsCounts[enteringNum] >= 0 { // вот здесь хотелось бы иметь более очевидную проверку на ISSET
			similarNumsCounts[enteringNum]++
		} else {
			similarNumsCounts[enteringNum] = 1
		}

		str = append(str, int(enteringNum))
	}

	fmt.Println(str)

	var max uint
	for key := range similarNumsCounts {
		if similarNumsCounts[key] > max {
			max = similarNumsCounts[key]
		}
	}

	return int(max)
}
