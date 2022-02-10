package main

import (
	"fmt"
	"strings"
)

var romans = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}
var romansExcl = map[string]int{
	"CM": 900,
	"CD": 400,
	"XC": 90,
	"XL": 40,
	"IX": 9,
	"IV": 4,
}

/**
roman numbers to integer arabian converter
*/
func main() {
	result := romanToInt("MMMCMXCIX")
	fmt.Println(result)
}

func romanToInt(s string) int {
	result := 0

	for len(s) > 1 {
		seekRomanStr := string(s[0]) + string(s[1])

		arabian := romansExcl[seekRomanStr]
		if arabian == 0 {
			seekRomanStr = string(s[0])
			arabian = romans[seekRomanStr]
		}

		result += arabian
		s = strings.TrimPrefix(s, seekRomanStr)
	}

	for len(s) > 0 {
		seekRomanStr := string(s[0])

		arabian := romans[seekRomanStr]
		result += arabian
		s = strings.TrimPrefix(s, seekRomanStr)
	}

	return result
}
