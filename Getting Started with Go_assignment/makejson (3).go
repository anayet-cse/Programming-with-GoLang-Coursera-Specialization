package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var name string
var address string

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimRight(name, "\n")
	if err != nil {
	}

	fmt.Print("Enter your address: \n")
	address, err2 := reader.ReadString('\n')

	if err2 != nil {
	}
	address = strings.TrimRight(address, "\n")
	myMap := map[string]string{
		"name":    name,
		"address": address,
	}
	myJSON, err := json.Marshal(myMap)
	if err != nil {
	}
	fmt.Println("\n The JSON with your name and address is:")
	fmt.Println(string(myJSON))

}
