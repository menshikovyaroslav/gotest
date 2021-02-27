package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type P struct {
	fname string
	lname string
}

func main() {

	sli := []P{}

	var name string
	fmt.Printf("Enter the name of a text file: ")
	fmt.Scan(&name)

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//data := make([]byte, 64)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err != nil && err != io.EOF {
			break
		}

		var line = scanner.Text()
		words := strings.Split(line, " ")
		if len(words) != 2 {
			continue
		}

		fName := words[0]
		lName := words[1]

		var p = P{fname: fName, lname: lName}
		sli = append(sli, p)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for index, value := range sli {
		fmt.Printf("next slice item: %v %s %s \n", index, value.fname, value.lname)
	}

	var txt string
	fmt.Scan(&txt)
}
