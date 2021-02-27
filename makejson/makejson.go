package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	fmt.Println("Enter Name:")
	fmt.Scan(&name)

	var address string
	fmt.Println("Enter Address:")
	fmt.Scan(&address)

	var person = map[string]string{}
	person["name"] = name
	person["address"] = address

	b, _ := json.Marshal(person)

	fmt.Println("json: ", string(b))

	var txt string
	fmt.Scan(&txt)
}
