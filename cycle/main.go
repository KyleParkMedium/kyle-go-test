package main

import (
	"cycle/A"
	"cycle/B"

	// "cycle/A"
	"fmt"
)

func main() {

	// A.SayFromA()

	// a := A.A{}

	// // fmt.Println(a.SayFromA())
	// a.SayFromA()
	// a.SayFromB()

	// b := B.B{}

	// fmt.Println(b)
	// b.SayFromA()

	a := A.NewA()
	fmt.Println(a)

	b := B.NewB(a)
	fmt.Println(b)

}
