package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := newCounter(0)
	c.print()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.add(1)
			}
		}()

	}
	wg.Wait()
	c.print()
}

type counter struct {
	count int
	mutex *sync.Mutex
}

func newCounter(initVal int) *counter {
	var m sync.Mutex
	return &counter{
		count: initVal,
		mutex: &m,
	}
}

func (c *counter) add(val int) {
	c.mutex.Lock()
	c.count += val
	c.mutex.Unlock()
}

func (c *counter) print() {
	fmt.Println(c.count)
}
