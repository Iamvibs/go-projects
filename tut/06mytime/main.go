package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time")

	// Date in string is default format for formatting
	presentTime := time.Now()
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createDate := time.Date(2024, time.August, 11, 10, 10, 0, 0, time.UTC)
	fmt.Println(createDate.Format("02-01-2006 Monday"))
}
