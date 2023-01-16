package A

import (
	"cycle/B"
	"fmt"
)

type A struct{}

func NewA() *A {
	return &A{}
}

func (a *A) SayFromA() {
	fmt.Println("A say hello")
}

func (a *A) SayFromB() {
	tmp_b := B.NewB(a)
	tmp_b.SayFromB()
}
