package main

import (
	"encoding/json"
	"fmt"
)

// we can add alias to structs keys with “
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // - means it will not be included when consuming...
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	vibsCourses := []course{
		{"ReactJS", 299, "YouTube", "qwerty123", []string{"js", "web-dev"}},
		{"Golang", 399, "YouTube", "abcd123", []string{"go", "backend"}},
		{"Solidity", 599, "YouTube", "mnb123", nil},
	}

	// pack this data as json..

	finalJson, err := json.MarshalIndent(vibsCourses, "", "\t") // \t is for clean json output (tab) and prefix is well...
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonData := []byte(`
	{
		"coursename": "ReactJS",
		"Price": 299,
		"website": "YouTube",
		"tags": ["js","web-dev"]
	}
	`)

	var vibsCourse course

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonData, &vibsCourse)
		fmt.Printf("%#v\n", vibsCourse)
	} else {
		fmt.Println("Invalid")
	}

	// some cases where we just want to add data to key value

	var myData map[string]interface{} // this is mapping... and interface if we dont know what is coming like array....
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%#v\n", vibsCourse)

	for k, v := range myData {
		fmt.Printf("Key is %v and value is %v\n", k, v)
	}
}
