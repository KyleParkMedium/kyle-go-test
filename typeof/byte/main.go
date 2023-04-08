package main

import (
	"fmt"
)

type A struct {
	Name  int
	Value int
}

func main() {

	var data [1]byte
	data[0] = 1111

	fmt.Println(data)
	// a := A{1, 2}

	// b, err := json.Marshal(a)

	// fmt.Println(err)
	// fmt.Println(b)

	// aa := A{}

	// // var aa interface{}
	// // aa := make(interface{})

	// err = json.Unmarshal(b, &aa)
	// fmt.Println(aa)
}
