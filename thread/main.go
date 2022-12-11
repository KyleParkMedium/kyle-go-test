package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	fmt.Println("one")
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Increment2() {
	fmt.Println("two")
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	a := Counter{value: 5}
	fmt.Println(a)
	go a.Increment()
	go a.Increment2()

	fmt.Println(a)
	// a.Increment()
	// fmt.Println(a)

}
