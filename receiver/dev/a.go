package dev

type Dev struct {
	Name int
}

func D() {
	C()

	aq := &Dev{Name: 100}
	aq.B()
}
