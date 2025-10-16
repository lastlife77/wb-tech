package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3}
	set2 := []int{2, 3, 4}

	fmt.Println(intersect(set1, set2))
}

func intersect(set1 []int, set2 []int) []int {

	res := []int{}

	for _, num1 := range set1 {
		for _, num2 := range set2 {
			if num1 == num2 {
				res = append(res, num1)
				break
			}
		}
	}

	return res
}
