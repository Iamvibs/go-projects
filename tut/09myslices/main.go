package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// to add things in slices
	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	// to get perfect values from slices and end value is non-enclusive so it will not be added
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	scores := make([]int, 4)

	// we cannot add more in this cause max len is 4 but we can append to it cause why not...
	scores[0] = 134
	scores[1] = 152
	scores[2] = 746
	scores[3] = 135

	scores = append(scores, 175, 555, 257, 467)

	// sort pkg to check or sort
	sort.Ints(scores)

	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))

	//remove value frm slices based on index
	var langs = []string{"Solidity", "GO", "JS", "Py", "java"}
	fmt.Println(langs)
	var index int = 2
	// basic remove
	langs = append(langs[:index], langs[index+1:]...)
	fmt.Println(langs)

}
