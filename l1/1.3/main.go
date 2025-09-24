package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	workersCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Ошибка: аргумент '%s' не является числом\n", os.Args[1])
		os.Exit(1)
	}
	ch := make(chan int)
	createPool(ch, workersCount)
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
