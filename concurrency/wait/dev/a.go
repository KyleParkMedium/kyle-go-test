package dev

import (
	"fmt"
	"sync"
)

func A(wg *sync.WaitGroup) {

	fmt.Println("hihi")

	wg.Done()
}
