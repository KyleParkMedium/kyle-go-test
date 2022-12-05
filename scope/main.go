package main

import "fmt"

func A(numA int) {

	numA += 1

	fmt.Println(numA)
}

func B(numB *int) {

	*numB += 1

	fmt.Println(&numB)
	fmt.Println(*numB)
	fmt.Println(numB)
}

func main() {

	numA := 5
	fmt.Println(numA)
	numB := 5

	A(numA)
	B(&numB)

	fmt.Println(numA)
	fmt.Println(numB)
}
