package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "\"hello\""

	fmt.Println(a)

	b := strconv.Quote("aa")

	c, _ := strconv.Unquote("AA")

	d := strconv.Itoa(123)
	e, _ := strconv.Atoi("123")
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
}
