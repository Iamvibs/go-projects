package main

import "fmt"

func main() {

	// LIFO when we use more than 1 defer
	// defer just execute loc at end of the function
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	fmt.Println("Hello")
	myDefer()
}

// hello, 01234, three, two, one

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
