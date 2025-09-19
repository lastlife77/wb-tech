package cncrmath

import (
	"sync"
	"testing"
)

func TestPow(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go Pow(3, ch, &wg)
	go func() {
		wg.Wait()
		close(ch)
	}()

	got := <-ch

	want := 9
	if got != want {
		t.Errorf("Pow(3) = %d, want %d", got, want)
	}
}

func TestPowArray(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int, 5)

	PowArray(arr, ch)

	got := 0
	for msg := range ch {
		got += msg
	}

	want := 220
	if got != want {
		t.Errorf("PowArray([2, 4, 6, 8, 10]) = %d, want %d", got, want)
	}
}
