package main

import (
	"fmt"
)

/**
16 Копилка: Мое кол-во денег в банке равно N. Каждый месяц я откладываю X денег, под процент Y.
Процент начисляется ежегодно. Нужно найти через какое кол-во лет я накоплю Z денег.
Входные данные: Четыре положительных числа
Выходные данные: фраза "Нужно ... лет" слово "лет" - "года" должно меняться
Sample Input: 2700000 50000 6 10000000
*/
var totalYears = 0
var yearsWords = map[int]string{0: "год", 1: "года", 2: "лет"}

func main() {
	var N int //мое кол-во денег в банке
	var X int // откладывание в месяц денег
	var Y int // процент под который откладываю
	var Z int // кол-во денег которое хочу

	fmt.Print("Enter current money in bank: ")
	scanLn, err := fmt.Scanln(&N)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	fmt.Print("Enter amount of money which you saving to bank monthly: ")
	scanLn, err = fmt.Scanln(&X)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	fmt.Print("Percent of deposit: ")
	scanLn, err = fmt.Scanln(&Y)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	fmt.Print("Which amount of money you want?: ")
	scanLn, err = fmt.Scanln(&Z)
	if err != nil {
		fmt.Print(scanLn, err)
	}

	totalYears := calcDeposit(N, X, Y, Z)

	correctYearWord := num2word(totalYears, yearsWords)

	fmt.Println("Нужно ", totalYears, " ", correctYearWord)
}

// SUMofPercents = (P*I)/100
// P - сумма привлеченных в депозит средств
// I - годовая процентная ставка
func calcDeposit(currentMoney int, moneyToSaving int, percent int, targetMoney int) int {
	totalYears++

	SUMofPercents := (currentMoney * percent * 12 / 12) / 100
	SUM := SUMofPercents + currentMoney + (moneyToSaving * 12)
	//fmt.Println("SUMofPercents - ", SUMofPercents)
	//fmt.Println("SUM - ", SUM)

	if SUM < targetMoney {
		SUM = calcDeposit(SUM, moneyToSaving, percent, targetMoney)
	}

	return totalYears
}

/**
 * Функция для определения окончания слова по числительному (1 день, 2 дня, 5 дней)
 *
 * @param num - число integer
 * @param words - массив с 3-мя словами
 */
func num2word(num int, words map[int]string) string {
	num = num % 100

	if num > 19 {
		num = num % 10
	}

	var word string

	if num == 1 {
		word = words[0]
	} else if num == 2 || num == 3 || num == 4 {
		word = words[1]
	} else {
		word = words[2]
	}

	return word
}
