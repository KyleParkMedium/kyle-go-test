package main

import (
	"fmt"
	"time"
)

func main() {

	a := time.Now()
	b := a.UTC()

	fmt.Println(a)
	fmt.Println(b)

	d := a.Format("6-01-0")

	fmt.Println(d)

}
