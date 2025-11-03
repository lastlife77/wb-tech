package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	s = delete(s, 4)
	fmt.Println(s)
}

func delete(s []int, i int) []int {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}
