package main

import (
	"fmt"
	"strconv"
)

func main() {

	// v := "0x1c"
	v := "1c"

	num, err := strconv.ParseUint(v, 16, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}

	v8 := uint8(num)
	fmt.Println(v8)

	// var a int64

	// a = 1

	// fmt.Println(a)

	// fmt.Println(a / 5000)

}
