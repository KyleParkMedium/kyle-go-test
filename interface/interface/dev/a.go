package dev

import "fmt"

type A interface {
	Q1() string
	// W1()

	// a string
}

type B struct {
	name int
	// QQ   A
}

type C interface {
	Q1() string
}

type D struct {
	name int
	QQ   C
}

func (d D) Q1() string {

	fmt.Println("Q1")

	return "hi"
}

func (b B) Q1() string {

	fmt.Println("Q1")

	return "hi"
}

func Q() {
	a := B{3}

	var qw C
	w := D{5, qw}
	fmt.Println(w)
	w.QQ.Q1()

	var b A

	b = a
	b.Q1()

	// fmt.Println("check")
	// fmt.Println(b)

	// // c := a.Q1()

	// b = a

	// c := b.Q1()
	// fmt.Println(c)
	// fmt.Println(a)
	// fmt.Println(b)
}
