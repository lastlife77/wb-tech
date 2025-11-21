package unpack

import (
	"errors"
	"strings"
	"testing"
)

// 5245344 225.3 ns/op 120 B/op 5 allocs/op
func BenchmarkDigitStrNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitStr(`a23bbbfn\312`)
	}
}

// 2797388 438.9 ns/op 332 B/op  19 allocs/op
func BenchmarkDigitStrOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		digitStr2(`a23bbbfn\312`)
	}
}

// digitStr2 is old version DigitStr without using strings.Builder
func digitStr2(s string) (string, error) {
	rs := []rune(s)

	if len(rs) < 2 {
		return s, nil
	}

	if isDigit(rs[0]) {
		return "", errors.New("string cannot start with a digit")
	}

	unpackS := string(rs[0])

	esc := false
	for i := 1; i < len(rs); i++ {
		if isEscape(rs[i]) {
			esc = true
			continue
		}
		if isDigit(rs[i]) && !esc {
			if i+1 < len(rs) && isDigit(rs[i+1]) {
				continue
			}
			c, j := getCount(rs, i)
			unpackS += strings.Repeat(string(rs[j]), c)
			continue
		}
		unpackS += string(rs[i])
		esc = false
	}

	return unpackS, nil
}
