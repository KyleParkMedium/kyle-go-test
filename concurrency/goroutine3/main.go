package main

import (
	"fmt"
	"time"
)

func A(ch chan int) {
	a := <-ch
	fmt.Println("A", a)
}
func B(ch chan int) {
	a := <-ch
	fmt.Println("B", a)
}

func main() {

	ch := make(chan int)

	for i := 0; i <= 10; i++ {
		go A(ch)
		go B(ch)
		ch <- i
	}

	time.Sleep(1000000)
}
