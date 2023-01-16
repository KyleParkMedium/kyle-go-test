package main

import (
	"fmt"
	"sync"
)

func A() {
	fmt.Println("hi")

}

func main() {

	// fmt.Println(len(recipients))
	var wg sync.WaitGroup
	// wg.Add(len(recipients))
	// wg.Add(3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			A()
			// defer wg.Done()

			wg.Done()
		}()

	}

	wg.Wait()

}
