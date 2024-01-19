package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
 *			STRUCT Person
 *	stores first and last names
 */

type Person struct {
	first string
	last  string
}

func main() {

	file, err := ioutil.ReadFile("Names.txt") // reads the file

	if err != nil {
		fmt.Println(err)
	}

	allNames := string(file)

	lines := strings.Split(allNames, "\n")

	data := make([]Person, 0, 1)
	for _, line := range lines {

		names := strings.Split(line, " ")
		data = append(data, Person{first: names[0], last: names[1]})

	}

	for _, name := range data {
		fmt.Println(name.first, name.last)
	}

}
