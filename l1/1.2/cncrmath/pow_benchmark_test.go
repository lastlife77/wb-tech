package cncrmath

import (
	"testing"
)

func BenchmarkPowArrayWithoutBuffer(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i + 1
	}

	for i := 0; i < b.N; i++ {
		ch := make(chan int)
		go PowArray(arr, ch)

		sum := 0
		for msg := range ch {
			sum += msg
		}
	}
}

func BenchmarkPowArrayWithBuffer(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i + 1
	}

	for i := 0; i < b.N; i++ {
		ch := make(chan int, len(arr))
		go PowArray(arr, ch)

		sum := 0
		for msg := range ch {
			sum += msg
		}
	}
}
