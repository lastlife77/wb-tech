package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)
	wg.Add(1)
	go gorutine(ch, &wg)
	time.Sleep((5 * time.Second))
	ch <- true
	wg.Wait()
}

func gorutine(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep((1 * time.Second))
			fmt.Println("Msg")
		}

	}
}
