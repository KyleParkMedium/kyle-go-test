package dev

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Println("Done")

}

func WaitGroup() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 100000)
	}

	wg.Done()
	wg.Wait()

}
