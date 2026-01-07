// Сombining done channels
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
		sig(1*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	var once sync.Once

	for _, ch := range channels {
		go func(c <-chan interface{}) {
			select {
			case <-c:
				once.Do(func() {
					close(out)
				})
				return
			case <-out:
				return
			}
		}(ch)
	}

	return out
}
