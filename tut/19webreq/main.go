package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("type of res is %T\n", res)

	defer res.Body.Close() // always close the connection manually

	databytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
