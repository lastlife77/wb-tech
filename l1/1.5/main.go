package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	app(5 * time.Second)
}

func app(timeout time.Duration) {
	ch := make(chan string)
	var wg sync.WaitGroup
	t := time.After(timeout)

	wg.Add(1)
	go writeMsgs(ch, t, &wg)

	wg.Add(1)
	go readMsgs(ch, &wg)

	wg.Wait()
}

func writeMsgs(ch chan<- string, timeout <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case ch <- randomMsg():
			fmt.Println("Отправлено сообщение.")
			time.Sleep(1 * time.Second)
		case <-timeout:
			fmt.Println("Время вышло! Завершение.")
			close(ch)
			return
		}
	}
}

func readMsgs(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("Сообщение:%s\n", msg)
	}
}

func randomMsg() string {
	return strconv.Itoa(rand.Intn(100))
}
