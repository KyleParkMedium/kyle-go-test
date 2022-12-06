package dev

import "fmt"

func (q *Dev) B() string {

	fmt.Println(q)
	q.Name -= 3
	fmt.Println(q)

	// q.Name
	return "ok"
}

func C() {
	fmt.Println("edddd")
}
