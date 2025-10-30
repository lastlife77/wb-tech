package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{
			s:   "english",
			exp: "hsilgne",
		},
		{
			s:   "русский",
			exp: "йикссур",
		},
		{
			s:   "%;№#@!",
			exp: "!@#№;%",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			act := rotate(test.s)
			if act != test.exp {
				t.Errorf("got:%v, want:%v", act, test.exp)
			}
		})
	}
}
