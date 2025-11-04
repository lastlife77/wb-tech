package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

func isUnique(s string) bool {
	m := make(map[rune]bool)

	s = strings.ToLower(s)

	for _, item := range s {
		if m[item] {
			return false
		}
		m[item] = true
	}

	return true
}
