package dev

import "fmt"

type Q struct {
	name int
}

func AA() {

	map_slice := map[int][]Q{}

	// map_slice := map[int][]int{}

	// map_slice[0][0] = int{11, 22}
	map_slice[0] = append(map_slice[0], Q{3})
	map_slice[0] = append(map_slice[0], Q{34})

	// map_slice[0] = append(map_slice[0], 1)
	// map_slice[0] = append(map_slice[0], 3)
	// map_slice[1] = []int{5}
	// map_slice[2] = []int{7}
	// map_slice[2] = append(map_slice[2], []int{9, 10}...)
	fmt.Println(map_slice)
}
