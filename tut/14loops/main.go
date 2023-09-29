package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday"}

	// basic for loop

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// golang for loop (works same)

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for each type of loop

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	value := 1

	for value < 10 {

		// goto

		if value == 5 {
			goto name
		}

		// break and continue

		if value == 5 {
			break
		}

		fmt.Println(value)
		value++
	}

	// this is lable which can be used with goto
name:
	fmt.Println("Vibhor")
}
