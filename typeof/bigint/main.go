package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	fmt.Println(a)

	a.SetInt64(654)
	fmt.Println(a)

	a.SetInt64(6541)
	fmt.Println(a)

	a.SetBit(a, 3, 1)
	fmt.Println(a)
}
