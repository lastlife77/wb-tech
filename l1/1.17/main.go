package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 8))
}

// accepts a sorted array
func binarySearch(arr []int, val int) int {
	length := len(arr)
	if length < 1 {
		return -1
	}
	guess := length / 2

	if arr[guess] == val {
		return guess
	}
	if arr[guess] > val {
		if guess-1 < 0 {
			return -1
		}
		return binarySearch(arr[:guess], val)
	} else {
		if guess+1 >= length {
			return -1
		}
		r := binarySearch(arr[guess+1:], val)
		if r == -1 {
			return r
		} else {
			return guess + 1 + r
		}
	}
}
