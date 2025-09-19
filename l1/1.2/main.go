package main

import (
	"fmt"

	"github.com/lastlife77/wb-tech/l1/1.2/cncrmath"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)

	cncrmath.PowArray(arr, ch)

	for msg := range ch {
		fmt.Println(msg)
	}
}
