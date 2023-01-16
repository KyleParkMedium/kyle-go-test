package main

import "fmt"

func A() {
	defer fmt.Println("hihi")

	defer fmt.Println("2")

	fmt.Println("1")
}
func main() {

	A()
}
