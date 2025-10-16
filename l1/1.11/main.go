package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3}
	set2 := []int{2, 3, 4}

	fmt.Println(intersect(set1, set2))
}

func intersect(set1 []int, set2 []int) []int {
	res := []int{}
	m := make(map[int]struct{})

	for _, num1 := range set1 {
		m[num1] = struct{}{}
	}

	for _, num2 := range set2 {
		if _, ok := m[num2]; ok {
			res = append(res, num2)
		}
	}

	return res
}
