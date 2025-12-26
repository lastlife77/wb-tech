// A program for searching for anagrams in a word slice.
package main

import (
	"fmt"

	"github.com/lastlife77/wb-tech/l2/2.11/anagram"
)

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"}
	anagrams := anagram.Search(words)
	fmt.Println(anagrams)
}
