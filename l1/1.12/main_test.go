package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestMakeSet(t *testing.T) {
	tests := []struct {
		arr []string
		exp []string
	}{
		{
			arr: []string{"cat", "cat", "dog", "cat", "tree"},
			exp: []string{"cat", "dog", "tree"},
		},
		{
			arr: []string{},
			exp: []string{},
		},
		{
			arr: []string{"cat"},
			exp: []string{"cat"},
		},
		{
			arr: []string{"cat", "cat"},
			exp: []string{"cat"},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			act := makeSet(test.arr)
			if !slices.Equal(act, test.exp) {
				t.Errorf("got: %v, want: %v", act, test.exp)
			}
		})
	}
}
