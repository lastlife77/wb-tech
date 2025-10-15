package main

import (
	"slices"
	"testing"
)

func TestGroupTmps(t *testing.T) {
	tests := []struct {
		tmps []float64
		exp  map[float64][]float64
	}{
		{
			tmps: []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			exp: map[float64][]float64{
				-20: {
					-25.4, -27.0, -21.0,
				},
				10: {
					13.0, 19.0, 15.5,
				},
				20: {
					24.5,
				},
				30: {
					32.5,
				},
			},
		},
		{
			tmps: []float64{1000.6, 1005.6, -102.5, -104.5, -5.4, -4.5, 4.5, 5.6},
			exp: map[float64][]float64{
				-100: {
					-102.5, -104.5,
				},
				-0.0001: {
					-5.4, -4.5,
				},
				0: {
					4.5, 5.6,
				},
				1000: {
					1000.6, 1005.6,
				},
			},
		},
	}

	for _, test := range tests {
		act := groupTmps(test.tmps)

		for key := range act {
			if !slices.Equal(act[key], test.exp[key]) {
				t.Errorf("Result was incorrect, got: %v, want: %v.", act[key], test.exp[key])
			}
		}
	}
}
