package dev

func A() string {

	p := &Dev{Name: 5}

	w := &Dev{Name: 10}

	e := &Dev{Name: 15}

	r := &Dev{Name: 20}

	p.B()
	w.B()
	e.B()
	r.B()

	// result := q.B()

	// c := a.Name

	// fmt.Println(result)

	// return result
	return "hi"
}
