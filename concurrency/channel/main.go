package main

import (
	"fmt"
	"sync"

	"channel/dev"
)

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch

	fmt.Println(n)
	wg.Done()
}
func main() {

	// var wg sync.WaitGroup
	// ch := make(chan int)

	// wg.Add(2)
	// go square(&wg, ch)
	// ch <- 9
	// wg.Wait()

	// dev.Dev()

	qqq := &dev.B{Name: 10}
	fmt.Println(qqq)

	fmt.Println(qqq.Z(4))
	fmt.Println(qqq.Dev2())
}
