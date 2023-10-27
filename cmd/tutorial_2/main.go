// 2. Data types, var vs const

package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	var intNum int16= 32000
	//int, int8, int16, int32 and int64 (and uint variants are there as well) are different variants of int. Also, compiler doesnt throw any errors during compile time if there is overflow of any value.
	// Also, int will default to 32 or 64 bytes(depending on system arch)
	fmt.Println(intNum);

	//floats are of two types - float32 and float64
	var floatNum float64=12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32=10.1
	var intNum32 int32=2
	var result float32=floatNum32+ float32(intNum32)
	fmt.Println(result)

	//NOTE: You cant perform operations with mixed types. You need to typecasting
	var intNum1 int=3
	var intNum2 int=2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string="Hello World" //or "Hello"+" "+"World"
	fmt.Println(myString)

	fmt.Println(len(myString))
	fmt.Println(len("γ")) //outputs 2 bytes (returns the number of bytes, rather than characters)
	fmt.Println(utf8.RuneCountInString("γ")) //alternative(returns the length of chars)

	var myRune rune='a'
	fmt.Println(myRune)

	var myBoolean=false
	fmt.Println(myBoolean)

	var intNum3 int //default value is 0
	fmt.Println(intNum3)
	
	var myVar="text" //the type is inferred
	fmt.Println(myVar)

	myVar2:="t1"
	fmt.Println(myVar2)
	
	var var1, var2 int=1,2 //OR var var1, var2=1,2 OR var1, var2:=1,2

	fmt.Println(var1, var2)

	//alternative to var
	const myConst string= "const value"
	fmt.Println(myConst)

	const pi float32=3.1415
}