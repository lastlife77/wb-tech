package main

import (
	"fmt"
	"math"
)

func main() {
	ts := []float64{
		-5.6, -6.4, 7.4, 5, 19.645, 15.5, 24.5, -21.0, 32.5, -15, -17, 105, 107, 124, -102,
	}
	groups := groupTmps(ts)
	fmt.Println(groups)
}

func groupTmps(ts []float64) map[float64][]float64 {
	groups := make(map[float64][]float64)

	for _, t := range ts {
		key := math.Trunc(t)
		if math.Abs(key) < 10 {
			if key >= 0 {
				key = 0
			} else {
				key = -0.0001 //аналог -0
			}
		} else {
			for k := 10.0; math.Abs(key/k) >= 1; k = k * 10 {
				key = key - float64(int(key)%10)
			}
		}

		groups[key] = append(groups[key], t)
	}

	return groups
}
