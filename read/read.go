package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	person := make(map[string]string)
	var name string
	var address string

	fmt.Printf("Name: ")
	fmt.Scan(&name)

	fmt.Printf("Address: ")
	fmt.Scan(&address)

	person["name"] = name
	person["address"] = address
	barr, _ := json.Marshal(person)
	fmt.Println(string(barr))

	var txt string
	fmt.Scan(&txt)
}
