package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go gorutine(&wg)
	wg.Wait()
}

func gorutine(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		if i == 5 {
			return
		}
	}
}
