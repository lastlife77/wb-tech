package main

import (
	"fmt"
)

func main() {
	fmt.Println(rotateText("snow dog sun"))
}

func rotateText(text string) string {
	runes := []rune(text)

	rotate(runes, 0, len(runes)-1)

	start := 0
	for i, r := range runes {
		if r == ' ' {
			rotate(runes, start, i-1)
			start = i + 1
		}
	}
	rotate(runes, start, len(runes)-1)

	return string(runes)
}

func rotate(runes []rune, left, right int) []rune {
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return runes
}
