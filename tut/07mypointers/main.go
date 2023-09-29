package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	myNumber := 20

	// & is reference and stores actual memory location
	var ptr = &myNumber

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("value of pointer is  ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is ", *ptr)

	// its nil
	var ptrrr *int
	fmt.Println(*ptrrr)

}
