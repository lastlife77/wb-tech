package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(makeSet(arr))
}

func makeSet(arr []string) []string {
	res := []string{}
	m := make(map[string]struct{})

	for _, s := range arr {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			res = append(res, s)
		}
	}

	return res
}
