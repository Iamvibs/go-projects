package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO")

	myCh := make(chan int, 1) // 1 is buffer
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-myCh

		fmt.Println(val)
		fmt.Println(isChanOpen)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
