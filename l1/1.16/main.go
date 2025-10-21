package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 6, 5, 7, 1, 2, 4}
	fmt.Println(quickSortStart(arr))
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var pivot int
		arr, pivot = partition(arr, low, high)
		arr = quickSort(arr, low, pivot-1)
		arr = quickSort(arr, pivot+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]

	return arr, i
}
