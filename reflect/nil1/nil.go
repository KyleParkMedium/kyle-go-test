package nil1

import (
	"fmt"
	"reflect"
)

type A struct {
	name int
}

func IsZero() {

	a := map[string]A{}

	a["A"] = A{1}
	a["B"] = A{1}

	// fmt.Println(a["C"])

	fmt.Println(len(a))
	// if *a["A"] == nil {
	// 	fmt.Println("correct")
	// }

	// w := reflect.TypeOf(a["A"])
	w := reflect.ValueOf(a["C"]).IsZero()
	// z := reflect.StructOf(a["A"].name).Size()
	fmt.Println(w)
}
