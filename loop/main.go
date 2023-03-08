package main

import "fmt"

func main() {
	a := make(map[string]int, 3)

	a["A"] = 1
	a["B"] = 2
	a["C"] = 3

	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(i)
		fmt.Println(a[i])
	}

}
