package main

import (
	"concurrency/testcode"
	"fmt"
)

func main() {

	data := 0

	data++

	if data == 0 {
		fmt.Println("0임")
	}

	testcode.Q()
}
