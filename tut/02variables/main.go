package main

import "fmt"

// Globle declaration

const LoginToken string = "qwertyuio"

func main() {
	// Explicite declaration

	var username string = "Vibhor"
	fmt.Println(username)
	fmt.Printf("Variables is of type %T \n", username)

	var number uint8 = 123
	fmt.Println(number)
	fmt.Printf("Variables is of type %T \n", number)

	var floting float32 = 123.12134834
	fmt.Println(floting)
	fmt.Printf("Variables is of type %T \n", floting)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type %T \n", isLoggedIn)

	// Inplicite declaration

	var user = "Mahajan"
	fmt.Println(user)

	password := 123456789
	fmt.Println(password)

	fmt.Println(LoginToken)

	// Default values

	var newUser string
	var newPass int

	fmt.Println(newPass)
	fmt.Println(newUser)
}
