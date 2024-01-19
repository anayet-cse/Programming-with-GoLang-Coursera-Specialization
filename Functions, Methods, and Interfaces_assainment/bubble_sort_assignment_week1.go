package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter 10 or less comma separated integers")
	userInput, _ := scanner.ReadString('\n')
	// remove any new line chars from user input
	userInput = strings.TrimSuffix(userInput, "\r\n")
	// trim all spaces from user input, if any
	userInput = strings.ReplaceAll(userInput, " ", "")
	userInputSplit := strings.Split(userInput, ",")

	// create the slice with values from user
	var slice []int
	for _, val := range userInputSplit {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			fmt.Printf("Error in converting val (type:%T) to int", val)
		}
		slice = append(slice, valInt)
	}

	fmt.Println("UNSORTED SLICE from USER")
	fmt.Println(slice)

	BubbleSort(slice)

	// print after sort
	fmt.Println("SORTED SLICE")
	for _, val := range slice {
		fmt.Print(val, " ")
	}
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				swap(slice, j, j+1)
			}
		}
	}
}

func swap(slice []int, i int, j int) {
	temp := slice[i]
	slice[i] = slice[j]
	slice[j] = temp
}
