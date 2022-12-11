package main

import (
	"concurrency/testcode"
	wg "concurrency/waitgroup"
	"fmt"
)

func main() {

	data := 0

	data++

	if data == 0 {
		fmt.Println("0임")
	}

	testcode.Q()

	// wg.WaitRange2()
	// wg.Memory()
	// wg.Wait3()
	// wg.RWMutex()
	wg.Cond()
}
