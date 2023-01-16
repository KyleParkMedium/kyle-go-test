package dev

import (
	"fmt"
	"reflect"
)

type B struct {
	Name string
	Age  int
	Num  int

	Array map[string]interface{}
}

func Basic() {

	// a := B{"kyle", 20, 5}
	b := B{}
	b.Name = "andy"

	// fmt.Println(a)
	fmt.Println(b)

	fmt.Println(reflect.TypeOf(b.Array))

}
