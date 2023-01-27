package main

import (
	"fmt"
)

func A(numA int) {

	numA += 1

	fmt.Println(numA)
}

func B(numB *int) {

	*numB += 1

	fmt.Println(&numB)
	fmt.Println(*numB)
	fmt.Println(numB)
}

type QQQ struct {
	a bool
}

func C() *QQQ {

	return new(QQQ)
}

// func C() *int {
// 	a := 1
// 	return &a
// }

func main() {
	a := 1
	b := &a
	fmt.Println(*b)

	// a := C()
	// fmt.Println(a)
	// fmt.Println(reflect.TypeOf(a))

	// numA := 5
	// fmt.Println(numA)
	// numB := 5

	// A(numA)
	// B(&numB)

	// fmt.Println(numA)
	// fmt.Println(numB)

	// a := 1
	// b := &a
	// c := &b

	// fmt.Println(a)
	// fmt.Println(&a)
	// fmt.Println(b)
	// fmt.Println(&b)
	// fmt.Println(&c)
}
