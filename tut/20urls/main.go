package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=react&paymentid=dmabn134"

func main() {
	fmt.Println("Welcome to URL")

	// parsing
	result, _ := url.Parse(myurl)

	// we can get everythings like this eg: port and rawQuery, path
	fmt.Println(result.Scheme)

	// this comes in key value pairs
	qparams := result.Query()

	fmt.Println(qparams["coursename"])

	// we can loop through it..

	for _, val := range qparams {
		fmt.Println(val)
	}

	// to create URL from data always use &

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/css",
		RawPath: "user=vibhor",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
