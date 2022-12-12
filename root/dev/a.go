package dev

import (
	"test"
	"test/dev/tt"
)

type B struct {
	Name string
	Age  string
}

func Q() {
	a := test.A{"1", "2"}

	b := tt.W()
	// fmt.Println(a)
}
