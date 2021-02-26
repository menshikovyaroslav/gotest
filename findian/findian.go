package main

import (
	"fmt"
	"strings"
)

func main() {

	var name string
	fmt.Println("Enter a string")
	fmt.Scan(&name)

	if len(name) < 3 {
		fmt.Println("Not Found!")
		return
	}

	name = strings.ToLower(name)

	var contains bool

	for i := 1; i < len(name)-1; i++ {
		if name[i] == 'a' {
			contains = true
		}
	}

	if name[0] == 'i' && name[len(name)-1] == 'n' && contains {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

	/* var name float32
	fmt.Println("Enter a floating point number")
	fmt.Scan(&name)
	var i int = int(name)
	fmt.Println("i = ", i)
	fmt.Scan(&name)
	*/

	fmt.Scan(&name)

	//	s := strings.Replace("ianianian", "ni", "in", 2)
	//	fmt.Println(s)

}
