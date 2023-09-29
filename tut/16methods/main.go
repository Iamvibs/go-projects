package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")

	Vibhor := User{"Vibhor", "vm11082001@gmail.com", true, 21}

	fmt.Println(Vibhor.Email)
	fmt.Printf("Vibhor details are %+v\n", Vibhor)

	Vibhor.GetStatus()
	Vibhor.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

// this just pass copy
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New email is ", u.Email)
}
