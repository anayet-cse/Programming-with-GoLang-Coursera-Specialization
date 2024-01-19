/*
A race condition is a situation that may occur inside a critical section. 
This happens when the result of multiple thread execution in critical section differs according to the order in which the threads execute.

Race conditions in critical sections can be avoided if the critical section is treated as an atomic instruction. 
Also, proper thread synchronization using locks or atomic variables can prevent race conditions.
*/

package main

import ("fmt"
        "time"
)
func increment(arr []int) {
    for i:=0 ; i<len(arr) ; i++ {
        arr[i]++
    }
}
func multiply(arr []int) {
    for i:=0 ; i<len(arr) ; i++ {
        arr[i] = arr[i] * 2
    }
}

func main() {
    arr := make([]int,1000)
    go increment(arr)
    go multiply(arr)
    time.Sleep(time.Second*4)
    
    sum := 0
    for i:=0 ; i<len(arr) ; i++ {
        sum =sum + arr[i]
    }
    fmt.Println(sum)
}