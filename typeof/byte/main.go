package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name  int
	Value int
}

func main() {

	a := A{1, 2}

	b, err := json.Marshal(a)

	fmt.Println(err)
	fmt.Println(b)

	aa := A{}

	// var aa interface{}
	// aa := make(interface{})

	err = json.Unmarshal(b, &aa)
	fmt.Println(aa)
}
