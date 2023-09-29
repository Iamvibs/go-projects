package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions")

	res := adder(69, 69)
	fmt.Println(res)

	proRes := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(proRes)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) int {
	result := 0

	// we dont care about index here so we dashed it
	for _, val := range values {
		result += val
	}
	return result
}
