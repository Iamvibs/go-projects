package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	Vibhor := User{"Vibhor", "vm11082001@gmail.com", true, 21}

	fmt.Println(Vibhor.Email)
	fmt.Printf("Vibhor details are %+v\n", Vibhor)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
