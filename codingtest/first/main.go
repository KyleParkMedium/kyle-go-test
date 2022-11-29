package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64 = 123123
	fmt.Println(n)

	s := strconv.FormatInt(n, 10)
	// s := strconv.Itoa(n)
	fmt.Println(s)
	fmt.Println(len(s))

	a := string(n)
	fmt.Println(a)
}
