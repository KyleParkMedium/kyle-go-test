package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	a := make(map[int8]int8)
	fmt.Println(a)

	b := []byte("AA")

	c := reflect.TypeOf(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(c)

	// a[b] = b

	z := "Abcd-1*ggggㅎㅎㅎ"
	for i := 0; i < len(z); i++ {
		fmt.Println(reflect.TypeOf(i))
		fmt.Println(reflect.TypeOf(z))
		fmt.Println(reflect.TypeOf(z[i]))
	}

	// z, _ := strconv.Atoi("AA")
	aa := "A"
	z, err := strconv.Atoi(aa)
	if err != nil {
		fmt.Errorf(err)
	}
	fmt.Println(z)

}
