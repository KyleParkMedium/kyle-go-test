package map3

import "fmt"

type AA struct {
	name int
}

func Map2() {

	a := make(map[string]interface{})

	a["AA"] = 11
	a["BB"] = 21

	for key, value := range a {

		fmt.Println(key)
		fmt.Println(value)
	}
}

func Map3() {

	// aa := make(map[string]map[string]AA)
	aa := make(map[string]map[string]AA)

	// aa := map[string]map[string]AA{name: 3}
	// aa["A"]["B"] = AA{}

	aa["A"] = make(map[string]AA)

	bb := AA{2}

	// aa["A"]["B"].name = 3
	aa["A"]["B"] = bb
	fmt.Println(aa)
}
