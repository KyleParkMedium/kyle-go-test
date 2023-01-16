package leakproofing

import (
	"fmt"
	"sync"
	"time"
)

func LeakProof() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})

		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)

			for {
				select {
				case s := <-strings:
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	// terminated := doWork(done, nil)

	terminated := doWork(done, nil)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated
	fmt.Println("Done.")
}

func Q() {
	fmt.Println("Q")
}
func W(wg *sync.WaitGroup) {
	// defer wg.Done()
	fmt.Println("W")
}
func Z(wg *sync.WaitGroup) {
	// defer wg.Done()
	fmt.Println("Z")
}

func R() {
	fmt.Println("R")
}

func E() {
	var wg sync.WaitGroup
	Q()

	// wg.Add(2)
	go Z(&wg)
	go W(&wg)

	// wg.Wait()

	// ch := make(chan interface{}, 1)

	// a := <-ch
	// fmt.Println(a)
	// // close(ch)
}
