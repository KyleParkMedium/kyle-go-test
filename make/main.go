package main

import "fmt"

func main() {

	a := make([]interface{}, 0)
	a = append(a, 1)
	a = append(a, "A")

	fmt.Println(a)
}
