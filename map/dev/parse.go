package dev

import (
	"fmt"
	"map/dev/q"
)

type D struct {
	DocType     string
	ID          string
	CreatedDate string
	UpdatedDate string
	Password    string

	// Name string
}

type C struct {
	DocType     string
	ID          string
	CreatedDate string
	UpdatedDate string
	Password    string
}

func W() {
	qwe := q.C{Password: "11"}
	fmt.Println(qwe)
	fmt.Println("hihi")
}

func E() {
	a := C{ID: "hi"}
	fmt.Println(a)
	a.Map()

	qwe := q.C{Password: "11"}
	qwe.Map()
}
