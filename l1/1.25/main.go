package main

import (
	"fmt"
	"time"
)

func main() {
	sleep1(3 * time.Second)
	sleep2(3 * time.Second)
	fmt.Println("End!")
}

func sleep1(d time.Duration) {
	done := make(chan struct{})

	go func() {
		start := time.Now()
		for time.Since(start) < d {

		}
		close(done)
	}()

	<-done
}

func sleep2(d time.Duration) {
	start := time.Now()
	for {
		if time.Since(start) >= d {
			break
		}
	}
}
