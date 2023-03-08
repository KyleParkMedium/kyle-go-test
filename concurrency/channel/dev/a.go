package dev

import "fmt"

type B struct {
	Name int
}

type A interface {
	// Hi() chan B
	// Hi2() <-chan B

	Z(a int) int
}

func (t *B) Z(a int) int {
	return 5
}

type Q struct {
	A
}

func Dev() {

	e := &B{Name: 10}
	fmt.Println(e)
	fmt.Println(e.Z(3))

	// fmt.Println(w)
}

type N struct {
	Age int
}

func (z *B) Dev2() {
	// q := &Q{A: z}

	// fmt.Println(q)
}
