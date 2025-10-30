package main

import "fmt"

func main() {
	fmt.Println(rotate("гла@вm5435ain***^&&ff"))
}

func rotate(s string) string {
	runes := []rune(s)

	i, j := 0, len(runes)-1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}
