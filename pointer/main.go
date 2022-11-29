package main

import "fmt"

func main() {

	a := 3
	b := &a
	var c *int
	c = &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
