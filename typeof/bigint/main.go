package main

import (
	"fmt"
	"math/big"
	"reflect"
)

type A struct {
	num big.Int
}

type B struct {
	num *big.Int
}

func main() {

	var a int64
	var b big.Int

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

	a = 5

	c := big.NewInt(5)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(*c))
	fmt.Println(reflect.TypeOf(&c))

	a := new(big.Int)
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	b := big.NewInt(1)
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(&b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(&b))

	q := A{*big.NewInt(1)}
	fmt.Println(q)
	fmt.Println(reflect.TypeOf(q))

	w := B{big.NewInt(1)}
	fmt.Println(w)
	fmt.Println(reflect.TypeOf(w))

	zz := big.NewInt(11)
	e := B{zz}
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))

	a.SetInt64(654)
	fmt.Println(a)

	a.SetInt64(6541)
	fmt.Println(a)

	a.SetBit(a, 3, 1)
	fmt.Println(a)

}
