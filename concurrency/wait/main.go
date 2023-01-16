package main

import (
	"sync"

	"wait/dev"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	for i := 0; i < 2; i++ {
		// go dev.A(&wg)
		go dev.A(&wg)

	}

	wg.Wait()

}
