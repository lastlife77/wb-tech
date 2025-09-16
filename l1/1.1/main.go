package main

import (
	"fmt"

	action "github.com/lastlife77/wb-tech/l1/1.1/Action"
	human "github.com/lastlife77/wb-tech/l1/1.1/Human"
)

func main() {
	h := human.New("Joseph", 17, true)
	a := action.New("Joseph", 17, true)
	fmt.Println(h.Age)
	fmt.Println(h.IsMale())
	h.GetOlder()
	fmt.Println(h.Age)

	fmt.Println(a.Age)
	fmt.Println(a.IsMale())
	a.GetOlder()
	fmt.Println(a.Age)
}
