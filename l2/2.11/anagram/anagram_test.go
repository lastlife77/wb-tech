package anagram

import (
	"fmt"
	"slices"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		words []string
		exp   map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"},
			exp: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case: %v\n", i), func(t *testing.T) {
			act := Search(test.words)
			if !equalAnagrams(test.exp, act) {
				t.Fatalf("\nAct:\n%v\nExp:\n%v", act, test.exp)
			}
		})
	}
}

func TestEqualAnagrams(t *testing.T) {
	tests := []struct {
		a   map[string][]string
		b   map[string][]string
		exp bool
	}{
		{
			a: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
			b: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
			exp: true,
		},
		{
			a: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
			b: map[string][]string{
				"листок": {"листок", "слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
			exp: true,
		},
		{
			a: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
			b: map[string][]string{
				"листок": {"слиток", "столик"},
				"пятак":  {"пятак", "пятка", "тяпка"},
			},
			exp: false,
		},
		{
			a: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
			b: map[string][]string{
				"листок": {"слиток", "столик"},
			},
			exp: false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Case: %v\n", i), func(t *testing.T) {
			act := equalAnagrams(test.a, test.b)
			if act != test.exp {
				t.Fatalf("\nAct:\n%v\nExp:\n%v", act, test.exp)
			}
		})
	}
}

func equalAnagrams(exp, act map[string][]string) bool {
	if exp == nil && act == nil {
		return true
	}
	if exp == nil || act == nil {
		return false
	}

	if len(exp) != len(act) {
		return false
	}

	for k, expWords := range exp {
		actWords, ok := act[k]
		if !ok {
			return false
		}
		if !slices.Equal(expWords, actWords) {
			return false
		}
	}

	return true
}
