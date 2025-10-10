package main

import "fmt"

func main() {
	fmt.Println(setBit(5, 1, true))  //5(101) -> 7(111)
	fmt.Println(setBit(5, 2, false)) //5(101) -> 1(001)
}

// If val is true sets one in postion
// If val is false sets zero in postion
// Bits are counted from zero
func setBit(num int64, pos int, val bool) int64 {
	if val {
		return 1<<pos | num
	} else {
		return ^(1 << pos) & num
	}
}
