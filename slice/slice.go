package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	sli := []int{}

	for true {
		var value string
		fmt.Print("Enter an integer value: ")
		fmt.Scan(&value)

		if value == "X" {
			break
		} else {
			i, _ := strconv.Atoi(value)
			sli = append(sli, i)
		}

		sort.Ints(sli)
		fmt.Println(sli)
	}

	var txt string
	fmt.Scan(&txt)

}
