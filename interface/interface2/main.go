package main

import "fmt"

type A struct {
	name int
}

type D interface {
	E()
}

type C struct {
	name int

	B
	D

	O N
}

type N struct {
	age string
	M
}

type M interface {
	CC()
}

// func (aa N) CC() {
// 	fmt.Println("final")
// }

type B interface {
	Q()
	W()
}

// func Q() {
// 	fmt.Println("QQ")
// }

// func (b *B) Q() {
// 	fmt.Println("WW")
// }

// func (b *B) W() {
// 	fmt.Println("WW")
// }

// func (a *A) Q() {
// 	println("hi")
// }

func (qqq A) Q() {
	fmt.Println("check")
}

func (aq C) Q() {
	// fmt.Println("check")
	// fmt.Println(aq.name)
	aq.W()
	aq.E()
}

func (aq C) W() {
	fmt.Println("check")
	fmt.Println(aq.name)
}

func (aq C) E() {
	fmt.Println("check")
	fmt.Println(aq.name)
}

func main() {

	// aaa := A{}/kisielk/godepgraph/blob/master/example.png
	// aaa.Q()
	// c := B{}

	// c.W()

	// var bbb B = A{name: 11}

	// bbb.Q()

	// var qwe B

	// fmt.Println(qwe)

	rq := C{}
	rq.name = 123
	rq.Q()

	rq.O.CC()

	// fmt.Println(rq)
}
