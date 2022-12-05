package main

import (
	"fmt"
	dev "receiver/dev"
)

func main() {
	fmt.Println("dev")
	result := dev.A()
	fmt.Println(result)
	// A()
	dev.D()

}
