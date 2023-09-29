package main

import (
	"fmt"
	"log"
	"mongodbapi/router"
	"net/http"
)

func main() {
	fmt.Println("Welcome to MongoDB API")

	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
