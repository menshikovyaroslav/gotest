package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var num string
	sortedSlice := make([]int, 3)
	for {
		fmt.Printf("Enter number to continue, X to exit: ")
		fmt.Scan(&num)
		if num == "X" {
			break
		}
		n, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		sortedSlice = append(sortedSlice, n)
		sort.Ints(sortedSlice)
		fmt.Printf("%v\n", sortedSlice)
	}

}
