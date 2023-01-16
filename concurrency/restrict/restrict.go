package restrict

import (
	"bytes"
	"fmt"
	"sync"
)

func Adhoc() {
	// data := make([]int, 4)
	data := [5]int{1, 2, 3, 4, 5}
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		// close(handleData)
		for i := range data {
			handleData <- data[i]
		}

	}
	handleData := make(chan int)

	go loopData(handleData)
	// loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}

	// aaa := func() {
	// 	fmt.Println("aa")
	// 	// fmt.Println(aaa)
	// }
	// go aaa()
}

func Restrict() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5)

		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received : %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	consumer(results)

}

func Buffer() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())

	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")

	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	wg.Wait()
}
