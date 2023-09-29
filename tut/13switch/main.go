package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch")

	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("GO GO GO")
	case 2:
		fmt.Println("Move 2 spots")
	case 3:
		fmt.Println("Move 3 spots")
	case 4:
		fmt.Println("Move 4 spots")
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots and roll again")
	}
}
