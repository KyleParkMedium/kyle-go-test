package dev

import (
	"fmt"
	"time"
)

func A() {
	for i := 10; i < 15; i++ {
		fmt.Println(i)
	}
}

func B() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func Start() {
	go A()
	go B()

	time.Sleep(3 * time.Second)
}
