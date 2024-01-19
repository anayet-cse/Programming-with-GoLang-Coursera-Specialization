package main

import "fmt"

func main() {

	var fnum float64

	fmt.Print("Enter a floating point number : ")
	fmt.Scanf("%f", &fnum)

	var truncNum int = int(fnum)

	fmt.Print("The truncated number is : ", truncNum)
}
