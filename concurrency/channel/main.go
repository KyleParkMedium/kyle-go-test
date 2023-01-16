package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	fmt.Println(n)
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}
