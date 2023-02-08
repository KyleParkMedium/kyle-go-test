package dev

import "fmt"

type A interface {
	Add() string
	Get() int
}

type B interface {
	A
}

type C struct {
	B
}

func (a *C) Add() string {
	return "A"
}

func (a *C) Get() int {
	return 5
}

func Check() {
	b := &C{}
	fmt.Println(b.Add())
	fmt.Println(b.Get())

	c := &C{}
	d := c.Add()

	fmt.Println(d)

}
