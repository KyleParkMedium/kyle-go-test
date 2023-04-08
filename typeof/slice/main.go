package main

import "fmt"

type A struct {
	B []string
	C [][32]byte
}

func main() {
	a := A{}
	a.B = append(a.B, "A")
	a.B = append(a.B, "1")

	fmt.Println(a)
}
