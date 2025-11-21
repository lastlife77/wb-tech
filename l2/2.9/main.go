package main

import (
	"fmt"

	"github.com/lastlife77/wb-tech/l2/2.9/unpack"
)

func main() {
	res, err := unpack.DigitStr(`a20`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
