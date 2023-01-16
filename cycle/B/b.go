package B

import (
	"fmt"
)

type virA interface {
	SayFromA()
}

type B struct {
	A virA
}

func NewB(arg_a virA) *B {
	return &B{
		A: arg_a,
	}
}

func (b *B) SayFromB() {
	fmt.Println("B say hello")
}

func (b *B) SayFromA() {
	// tmp_a := A.NewB()
	// tmp_a.SayFromA()
	b.A.SayFromA()
}
