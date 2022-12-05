package main

import "fmt"

func A(msg int) {
	fmt.Println("Check")
	fmt.Println(msg)

}
func main() {

	a := make(chan int, 3)

	a <- 3
	fmt.Println(a)
	a <- 4
	fmt.Println(a)

	a <- 5
	fmt.Println(a)

	for {
		msg := <-a

		// b := A(msg)
		// fmt.Println("Ex")
		// fmt.Println(b)
		A(msg)

	}

}
