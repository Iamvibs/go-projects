package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web request verbs")
	GetRequest()
	PostRequest()
	PostFormRequest()
}

func GetRequest() {
	const myurl = "http://localhost:8000/get"

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("Status code is ", res.StatusCode)
	fmt.Println("Content length is ", res.ContentLength)

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PostRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"Lang": "Golang",
			"Learner": "Vibhor"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}

func PostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("fistName", "Vibhor")
	data.Add("lastName", "Mahajan")

	res, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}
