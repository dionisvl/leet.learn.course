package main

import (
	"fmt"
	"strings"
)

/**
Longest Common Prefix
*/
func main() {
	var strs = []string{"flower", "flow", "flght"}
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

func longestCommonPrefix(strs []string) string {
	currentPrefix := ""
	newPrefix := ""
	for i := 0; i < len(strs[0]); i++ {
		newPrefix += string(strs[0][i])
		for _, str := range strs {
			//fmt.Println(key)
			if strings.HasPrefix(str, newPrefix) {
				continue
			}
			return currentPrefix
		}
		currentPrefix = newPrefix

	}

	return currentPrefix
}
