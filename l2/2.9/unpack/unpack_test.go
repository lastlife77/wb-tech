package unpack

import (
	"errors"
	"fmt"
	"testing"
)

func TestUnpackStr(t *testing.T) {
	tests := []struct {
		str string
		exp string
		err error
	}{
		{
			str: "a4bc2d5e",
			exp: "aaaabccddddde",
			err: nil,
		},
		{
			str: "abcd",
			exp: "abcd",
			err: nil,
		},
		{
			str: "45",
			exp: "",
			err: errors.New("string cannot start with a digit"),
		},
		{
			str: "",
			exp: "",
			err: nil,
		},
		{
			str: "e",
			exp: "e",
			err: nil,
		},
		{
			str: `qwe\4\5`,
			exp: "qwe45",
			err: nil,
		},
		{
			str: `qwe\45`,
			exp: "qwe44444",
			err: nil,
		},
		{
			str: `a11bb`,
			exp: "aaaaaaaaaaabb",
			err: nil,
		},
		{
			str: `a2n\312`,
			exp: "aan333333333333",
			err: nil,
		},
		{
			str: `aaa\\\\2\43`,
			exp: "aaa2444",
			err: nil,
		},
		{
			str: `a20`,
			exp: "aaaaaaaaaaaaaaaaaaaa",
			err: nil,
		},
		{
			str: `a003`,
			exp: "aaa",
			err: nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			act, err := DigitStr(test.str)
			if err != nil {
				if test.err.Error() != err.Error() {
					t.Errorf("got: %v, want: %v", err, test.err)
				}
			}
			if act != test.exp {
				t.Errorf("got: %v, want: %v", act, test.exp)
			}
		})
	}
}
