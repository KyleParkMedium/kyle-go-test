package main

import "fmt"

func main() {

	a := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	b := [4]int{2, 2, 2, 2}
	for _, x := range a {
		A(x)
	}
	for _, y := range b {
		fmt.Println(y)
	}

}

func A(x [4]int) {

	for _, y := range x {
		fmt.Println(y)
	}

}
