package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		set1 []int
		set2 []int
		exp  []int
	}{
		{
			set1: []int{1, 2, 3},
			set2: []int{2, 3, 4},
			exp:  []int{2, 3},
		},
		{
			set1: []int{1, 2, 3},
			set2: []int{1, 2, 3},
			exp:  []int{1, 2, 3},
		},
		{
			set1: []int{},
			set2: []int{1, 2, 3},
			exp:  []int{},
		},
		{
			set1: []int{1, 2, 3},
			set2: []int{},
			exp:  []int{},
		},
		{
			set1: []int{},
			set2: []int{},
			exp:  []int{},
		},
		{
			set1: []int{51, 56, 7, 8, 90},
			set2: []int{7, 90, 4, 34, 5, 6},
			exp:  []int{7, 90},
		},
		{
			set1: []int{3, 4, 7, 90},
			set2: []int{5, 67, 8, 34, 6},
			exp:  []int{},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			act := intersect(test.set1, test.set2)
			if !reflect.DeepEqual(act, test.exp) {
				t.Errorf("got %v, want %v", act, test.exp)
			}
		})
	}
}
