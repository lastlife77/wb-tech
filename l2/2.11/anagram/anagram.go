// Package anagram provides utilities for finding anagrams in a slice of words.
package anagram

import (
	"slices"
)

// Search groups anagrams from the given words.
// It returns a map where each key is the first encountered word from an anagram group,
// and the value is a slice containing all words from that group (including the key).
// The groups are sorted alphabetically, and each group's words are also sorted,
// excluding the key which appears first in its original form.
func Search(words []string) map[string][]string {
	anagramsMap := make(map[string][]string, len(words)/2)

	// n = len(words)
	// O(n) * O(m*log(m)) = O(n*m*log(m))
	for _, word := range words {
		letters := []rune(word)
		// m = average len(letters)
		// O(m*log(m))
		slices.Sort(letters)
		anagram := string(letters)
		anagramsMap[anagram] = append(anagramsMap[anagram], word)
	}

	sortedAnagramsMap := make(map[string][]string, len(anagramsMap))

	// k = len(anagrams)
	// O(k) * O(v*log(v)) = O(k*v*log(v))
	// k*v = n = > O(n*log(n)) - worst case
	for k, v := range anagramsMap {
		if len(v) < 2 {
			delete(anagramsMap, k)
			continue
		}
		// v = average len(letters)
		// O(v*log(v))
		slices.Sort(v)
		sortedAnagramsMap[v[0]] = v
		delete(anagramsMap, k)
	}

	// O = O(n*m*log(m)) + O(n*log(n)) = O(n*(m*log(m) + log(n)))
	// m*log(m) > log(n) => O(n*(m*log(m))
	return sortedAnagramsMap
}
