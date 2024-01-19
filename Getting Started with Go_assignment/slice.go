package main

import (
   "fmt"
   "strconv"
   "sort"
)

func main () {
   
   sli := make([]int, 0, 3)
   a := ""
   
   fmt.Println("Enter an integer to append to slice or enter X to print slice and exit program")
   fmt.Scanln(&a)
   
   for a != "X" {
      b, err := strconv.Atoi(a)
	  if err != nil {
	     fmt.Println("Error")
		 } else {
		 sli = append(sli, b)
		 sort.Ints(sli)
		 fmt.Println(sli)
		 fmt.Println("Enter an integer to append to slice or enter X to print slice and exit program")
		 fmt.Scanln(&a)
		 }
	}
	fmt.Println(sli)
}