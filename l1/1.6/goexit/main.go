package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go goroutine(&wg)

	wg.Wait()
}

func goroutine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; ; i++ {
		if i == 5 {
			runtime.Goexit()
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Msg")
	}
}
