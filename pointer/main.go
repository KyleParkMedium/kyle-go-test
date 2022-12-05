package main

import "fmt"

type A struct {
	Name int
	Age  int
}

func main() {

	a := 3
	b := &a
	var c *int
	c = &a

	var d *A = &A{1, 2}
	var e A = A{3, 4}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(*c)
	fmt.Println(&c)
	fmt.Println(d)
	fmt.Println(&d)
	fmt.Println(e)
	fmt.Println(&e)
}
