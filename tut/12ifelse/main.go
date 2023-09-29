package main

import "fmt"

func main() {
	fmt.Println("Welcome to ifelse")

	// normal ifelse statements works
	// if cond {} else if cond {} else {}

	// new in golang

	if num := 11; num < 10 {
		fmt.Println("Number is smaller")
	} else {
		fmt.Println("Number is greater")
	}

	// for err handling
	/*
		if err != nil {
			condition
		}
	*/
}
