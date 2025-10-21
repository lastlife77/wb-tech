package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		arr []int
		exp []int
	}{
		{
			arr: []int{},
			exp: []int{},
		},
		{
			arr: []int{1},
			exp: []int{1},
		},
		{
			arr: []int{3, 2},
			exp: []int{2, 3},
		},
		{
			arr: []int{1, 2, 5, 7, 4},
			exp: []int{1, 2, 4, 5, 7},
		},
		{
			arr: []int{9, 8, 7, 6},
			exp: []int{6, 7, 8, 9},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			act := quickSortStart(test.arr)
			if !slices.Equal(act, test.exp) {
				t.Errorf("got: %v, want: %v", act, test.exp)
			}
		})
	}
}
