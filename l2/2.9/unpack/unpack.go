// Package unpack package provides unpacking of various sequences.
package unpack

import (
	"errors"
	"strings"
)

// DigitStr unpacks a string like "a4bc2d5e" into "aaaabccddddde".
// Supports escape sequences like '\4',
// which will be interpreted as a character rather than a number.
// If S starts with a digit, an error will be returned.
// Several consecutive digits are perceived as one number:
// "a11" = "aaaaaaaaaaa".
func DigitStr(s string) (string, error) {
	rs := []rune(s)

	if len(rs) < 2 {
		return s, nil
	}

	if isDigit(rs[0]) {
		return "", errors.New("string cannot start with a digit")
	}

	var builder strings.Builder
	builder.WriteRune(rs[0])

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
			builder.WriteString(strings.Repeat(string(rs[j]), c))
			continue
		}
		builder.WriteRune(rs[i])
		esc = false
	}

	return builder.String(), nil
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isEscape(r rune) bool {
	return r == '\\'
}

func getCount(rs []rune, i int) (int, int) {
	count := int(rs[i] - '0')
	i--
	k := 10
	for i > 0 && isDigit(rs[i]) && !isEscape(rs[i-1]) {
		count += int(rs[i]-'0') * k
		i--
		k *= 10
	}

	return count - 1, i
}
