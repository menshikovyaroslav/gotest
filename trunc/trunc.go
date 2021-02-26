package main

import (
	"fmt"
)

func main() {

	var name float32
	fmt.Println("Enter a floating point number")
	fmt.Scan(&name)
	var i int = int(name)
	fmt.Println("i = ", i)
	fmt.Scan(&name)
}
