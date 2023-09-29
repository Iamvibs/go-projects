package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	lang := make(map[string]string)

	lang["JS"] = "javascript"
	lang["SOL"] = "solidity"
	lang["PY"] = "python"
	lang["GO"] = "golang"

	fmt.Println(lang)
	fmt.Println(lang["SOL"])

	// to delete
	delete(lang, "JS")
	fmt.Println(lang)

	for _, value := range lang {
		fmt.Printf(value)
	}
}
