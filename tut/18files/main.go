package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files")

	content := "Name: Vibhor Mahajan"

	file, err := os.Create("./info.txt")
	checkNilErr(err)

	len, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println(len)

	defer file.Close()

	readFile("./info.txt")
}

func readFile(file string) {
	data, err := os.ReadFile(file)
	checkNilErr(err)
	fmt.Println(string(data))
}

// this is common practice
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
