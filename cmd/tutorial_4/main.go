package main

import "fmt"

//4 points about Arrays in Golang:
// 1) Fixed length
// 2) Stores elements of same type
// 3) Indexable
// 4) Contiguous in memory
// 5) Default values in the array depends on the default value of the element types on which the array is based. In this case, for int32, it is 0

func main(){
	var intArr [3]int32
	intArr[1]=132
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) //access 1st and 2nd element
}