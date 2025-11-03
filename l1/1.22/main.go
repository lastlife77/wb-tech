package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("1000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("5000000000000000000000", 10)
	fmt.Printf("%v * %v = %v\n", a, b, new(big.Int).Mul(a, b))
	fmt.Printf("%v / %v = %v\n", a, b, new(big.Int).Div(a, b))
	fmt.Printf("%v + %v = %v\n", a, b, new(big.Int).Add(a, b))
	fmt.Printf("%v - %v = %v\n", a, b, new(big.Int).Sub(a, b))
}
