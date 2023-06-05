package main

import (
	"fmt"
	"reflect"
)

func A(q ...string) {
	fmt.Println(q)

	for _, v := range q {
		fmt.Println(v)
	}
}
{"type": "ethereum", "signature": mintTxDto.signature}, {}
func sum(numbers ...int) int {
	fmt.Println(numbers)
	fmt.Println(reflect.TypeOf(numbers))
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	_ = sum(1, 2, 3, 4)

	// q := []string{"a", "B"}
	// A("a", "B")
	// A(q)
}
