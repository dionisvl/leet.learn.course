package main

import (
	"fmt"
)

// 242. Valid Anagram
func main() {
	TestIsEven()
}

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(s) != len(t) {
		return false
	}
	type LetterCount struct {
		Letter rune // Строчная английская буква, представленная в виде Unicode кода
		Count  int  // Количество повторений данной буквы, по умолчанию равно нулю
	}
	var lowercaseAlphabet1 [26]LetterCount
	var lowercaseAlphabet2 [26]LetterCount
	for _, ch := range s { // перебираем все символы в строке
		lowercaseAlphabet1[ch-'a'].Letter = ch // сохраняем её в массиве структур под индексом её номера в алфавите
		lowercaseAlphabet1[ch-'a'].Count++     // увеличиваем значение Count для этой буквы на 1
	}
	for _, ch := range t { // перебираем все символы в строке
		lowercaseAlphabet2[ch-'a'].Letter = ch // сохраняем её в массиве структур под индексом её номера в алфавите
		lowercaseAlphabet2[ch-'a'].Count++     // увеличиваем значение Count для этой буквы на 1
	}
	return lowercaseAlphabet1 == lowercaseAlphabet2
}

func TestIsEven() {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{"1", "amma", "mama", true},
		{"2", "", "mama", false},
		{"4", "aacc", "cacc", false},
		{"5", "anagram", "nagaram", true},
		{"6", "a", "a", true},
	}

	for _, tc := range testCases {
		result := isAnagram(tc.s, tc.t)
		if result != tc.expected {
			fmt.Println("expected %v but got %v in case%v", tc.expected, result, tc.name)
		}
	}
}
