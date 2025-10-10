package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go counter(nums)
	go squarer(nums, squares)
	printer(squares)
}

func counter(nums chan<- int) {
	for i := 0; i < 100; i++ {
		nums <- rand.Intn(10)
		time.Sleep(100 * time.Millisecond)
	}
	close(nums)
}

func squarer(nums <-chan int, squares chan<- int) {
	for x := range nums {
		squares <- x * x
	}
	close(squares)
}

func printer(squares <-chan int) {
	for x := range squares {
		fmt.Println(x)
	}
}
