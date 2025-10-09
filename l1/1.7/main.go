package main

import (
	"fmt"
	"sync"
)

func main() {
	m := NewMySyncMap()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			m.Write("test", i)
			fmt.Printf("Write: %d \n", i)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			val, success := m.Read("test")
			fmt.Printf("Read: %v %v \n", val, success)
		}()
	}
	wg.Wait()
}

type MySyncMap struct {
	mx sync.RWMutex
	m  map[string]int
}

func NewMySyncMap() *MySyncMap {
	return &MySyncMap{
		m: make(map[string]int),
	}
}

func (msm *MySyncMap) Read(key string) (int, bool) {
	msm.mx.RLock()
	defer msm.mx.RUnlock()
	val, ok := msm.m[key]
	return val, ok
}

func (msm *MySyncMap) Write(key string, value int) {
	msm.mx.Lock()
	defer msm.mx.Unlock()
	msm.m[key] = value
}
