package main

import (
	"fmt"
	"math/big"
)

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}

func main() {

	// var num1 int32
	var num1 int64
	// var loc *int
	var loc *int64

	num1 = 3
	loc = &num1
	fmt.Println(loc)

	fmt.Println(num1)
	fmt.Println(&num1)
	fmt.Println(num1 >> 3)
	// fmt.Println(~num1)

	num2 := 3<<3 - 1
	fmt.Println(num2)

	var a uint8
	a = 255
	a = a >> 12321
	fmt.Println(a)

	var b byte
	b = 255
	fmt.Println(b)

	tt256 := BigPow(2, 256)
	tt256m1 := new(big.Int).Sub(tt256, big.NewInt(1))

	fmt.Println(tt256)
	fmt.Println(tt256m1)

	r := big.NewInt(1)
	w := big.NewInt(10)
	e := big.NewInt(100)
	fmt.Println(r)
	fmt.Println(w)
	// z := big.Add(r, w)
	// z := big.Int.Add(&r, &w)

	x := new(big.Int)
	z := x.Add(x, x)

	e.Sub(r, w)
	fmt.Println(e)

	zz := r.Add(r, w)
	fmt.Println(zz)
	fmt.Println(x)
	fmt.Println(z)

}
