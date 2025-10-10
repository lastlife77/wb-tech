package main

import (
	"testing"
)

func TestSetBit(t *testing.T) {
	tests := []struct {
		num int64
		pos int
		val bool
		exp int64
	}{
		{num: 5, pos: 1, val: true, exp: 7},   //5(101) -> 7(111)
		{num: 5, pos: 2, val: false, exp: 1},  //5(101) -> 1(001)
		{num: 15, pos: 2, val: true, exp: 15}, //15(1111) -> 15(1111)
		{num: 2, pos: 3, val: false, exp: 2},  //2(0010) -> 2(0010)
		{num: 2, pos: 3, val: true, exp: 10},  //2(0010) -> 10(1010)
	}

	for _, test := range tests {
		act := setBit(test.num, test.pos, test.val)
		if act != test.exp {
			t.Errorf("Result was incorrect, got: %d, want: %d.", act, test.exp)
		}
	}
}
