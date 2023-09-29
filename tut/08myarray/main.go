package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	// normal decl
	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"

	fmt.Println("fruitlist is ", fruitList)

	// adv decl
	var vegList = [3]string{"potato", "spinach", "caouliflower"}

	fmt.Println("vegy list is ", vegList)

}
