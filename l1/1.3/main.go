package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)
	createPool(ch, 6)
	write(ch)
}

func createPool(ch chan int, workersCount int) {
	for i := 1; i <= workersCount; i++ {
		go worker(i, ch)
	}
}

func worker(id int, ch <-chan int) {
	for msg := range ch {
		fmt.Printf("Worker%d: %d \n", id, msg)
	}
}

func write(ch chan<- int) {
	for {
		ch <- rand.Intn(100)
		time.Sleep(1 * time.Second)
	}
}
