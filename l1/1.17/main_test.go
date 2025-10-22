package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		arr []int
		val int
		exp int
	}{
		{
			arr: []int{},
			val: 5,
			exp: -1,
		},
		{
			arr: []int{1},
			val: 2,
			exp: -1,
		},
		{
			arr: []int{1},
			val: 1,
			exp: 0,
		},
		{
			arr: []int{1, 5},
			val: 5,
			exp: 1,
		},
		{
			arr: []int{1, 5},
			val: 6,
			exp: -1,
		},
		{
			arr: []int{5, 8, 11, 45},
			val: 45,
			exp: 3,
		},
		{
			arr: []int{5, 43, 78, 100},
			val: 5,
			exp: 0,
		},
		{
			arr: []int{5, 43, 78, 100},
			val: 43,
			exp: 1,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			act := binarySearch(test.arr, test.val)
			if act != test.exp {
				t.Errorf("got: %v, want: %v", act, test.exp)
			}
		})
	}
}
