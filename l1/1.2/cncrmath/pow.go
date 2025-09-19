// Package cncrmath provides mathematical operations with concurrent execution.
package cncrmath

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

// PowArray writes the squares of all numbers in arr to ch concurrently
// and closes ch when done.
func PowArray[T number](arr []T, ch chan T) {
	var wg sync.WaitGroup
	for _, num := range arr {
		wg.Add(1)
		go Pow(num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
}

// Pow writes the square of num to ch and marks the WaitGroup as done.
func Pow[T number](num T, ch chan T, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- num * num
}
