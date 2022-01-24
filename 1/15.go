package main

import (
	"fmt"
	"strconv"
)

/**
15 Часы: Минутная стрелка повернулась с начала суток на N градусов.
Определите, сколько сейчас целых часов h и целых минут m. Слова "Час", "Минута" должны склоняться.
Входные данные: На вход программе подается целое число N, N > 0
Выходные данные: Московское время ... часов ... минут
Sample Input: 0 Sample Output: Московское время 0 часов 0 минут
*/
func main() {
	var degree int

	fmt.Print("Enter degree of minute arrow: ")
	scanLn, err := fmt.Scanln(&degree)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	answer := calcTime(degree)
	fmt.Println("result: ", answer)
}

func calcTime(degree int) string {
	minutes := degree / 60 // сколько всего минут
	minutesWithoutHours := degree % 60
	hours := minutes % 60

	return fmt.Sprintf(
		"Московское время %s часов %s минут",
		strconv.Itoa(hours),
		strconv.Itoa(minutesWithoutHours),
	)
}
