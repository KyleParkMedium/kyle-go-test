package main

import (
	"fmt"
	"interface5/dev"
)

type Q struct {
	AA dev.C
}

// func (a *Q) Add() string {
// 	return "B"
// }

// func (a *Q) Get() int {
// 	return 5101
// }

func (a *Q) ChCh() int {
	fmt.Println(a.AA.Add())
	return 5101
}

func main() {

	b := &Q{}
	fmt.Println(b)
	fmt.Println(b.ChCh())

	// dev.Check()
}
