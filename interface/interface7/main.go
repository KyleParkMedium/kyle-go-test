package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
	Age int
}

func AA(bytes []byte) interface{} {

	fmt.Println(bytes)

	a := A{}

	err := json.Unmarshal(bytes, &a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
	fmt.Println(a.Age)
	return a
}

func main() {

	a := A{Num: 3, Age: 10}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	b := AA(bytes)
	fmt.Println(b)
	fmt.Println(b)

}
