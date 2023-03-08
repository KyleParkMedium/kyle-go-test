package select1

import (
	"fmt"
	"sync"
	"time"
)

func Select() {

	// i := "hi"

	i := make(chan string)

	go func() {
		for {
			// fmt.Println("check")
			time.Sleep(5 * time.Second)
			i <- "1"
		}
	}()

	for {
		select {

		case a := <-i:
			fmt.Println(a)
			fmt.Println("good")

		case a := <-i:
			fmt.Println(a)
			fmt.Println("good")

		default:

		}

		// scheduling()

	}
}

func scheduling() {
	fmt.Println("sche")
	//do something
}

func Example(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("exm")
	// wg.Done()
}

func Routine() {
	a := make([]int, 10)
	var wg sync.WaitGroup
	// wg.Add(len(a))
	for _, _ = range a {
		wg.Add(1)
		go Example(&wg)
	}
	wg.Wait()
}

type Msg struct {
	msg int
	err error
}

func Example2(i *int, aa chan int) {

	defer AB()
	// *i++

	aa <- 3
	// aa <- &Msg{msg: 3}
}

func Example3(aa chan *Msg) {

	defer AB()
	// *i++

	// aa <- 3
	// aa <- &Msg{msg: 3}

}

func AB() {

}

func CH() {

	a := make([]int, 10)
	q := make(chan int)
	w := make(chan *Msg, 1)

	i := 0

	// q <- &Msg{msg: 3}
	// fmt.Println(q)

	for _, _ = range a {
		go Example2(&i, q)
		go Example3(w)
	}

	for {

		if i == 10 {
			fmt.Println("here")
			return
		}

		select {

		case b := <-q:

			fmt.Println(b)
			i++
			// switch {
			// case b.err == io.EOF:
			// 	fmt.Println("good")
			// 	return
			// default:
			// 	fmt.Println(b)
			// }

		default:

		}

	}
}
