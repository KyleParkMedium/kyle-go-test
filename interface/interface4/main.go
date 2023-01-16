package main

import "fmt"

type A interface {
	Q()
}

type Folder struct {
	B []A
}

type G struct {
	// check string
}

func (g *G) Q() {

}

func (z *Folder) Q() {

}

func (g *G) W() {

}

func main() {

	q := &G{}
	// a := &Folder{B: []A{q, q}}

	a := &Folder{}

	fmt.Println(a)

	var w []A
	// var e A
	var n A = &G{}

	// w = append(w, e)
	w = append(w, n)
	w = append(w, n)
	w = append(w, n)
	w = append(w, n)
	w = append(w, n)
	w = append(w, n)

	a.B = w

	b := &Folder{B: []A{a, q, q}}

	fmt.Println(b)

}
