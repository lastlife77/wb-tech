package main

import (
	"fmt"
	"testing"
)

func TestRotateText(t *testing.T) {
	tests := []struct {
		text string
		exp  string
	}{
		{
			text: "snow dog sun",
			exp:  "sun dog snow",
		},
		{
			text: "123 45 67",
			exp:  "67 45 123",
		},
		{
			text: "любовь теряет разум",
			exp:  "разум теряет любовь",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			act := rotateText(test.text)
			if act != test.exp {
				t.Errorf("Got: %v, want: %v", act, test.exp)
			}
		})
	}
}
