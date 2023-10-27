// 5. Strings, bytes and runes

package main

import (
	"fmt"
	"reflect"
	"strings"
)

//NOTE: ascii uses 7 bits while utf-8 uses 8 bits. In UTF-8, upto ascii code point 127(starting from 0), 1 byte is used to represent it in utf-8, with the MSB bit set to 0. Hence, only 0 to 127 numbered characters can be stored using utf-8 as the 1 bit is always 0
//REFERENCE: https://www.youtube.com/watch?v=8uiZC0l4Ajw Timestamp: 28:18 and 29:05

//In Golang, String is basically array of bytes

func main(){
	var myString="résumé"
	var indexed=myString[0] //returns uint8 value which is basically the utf-8 value corresponding to the character
	indexed=myString[1] //gives only the first half(first byte value) of é (returns 195)

	fmt.Println(indexed, reflect.TypeOf(indexed)) 
	//OR
	fmt.Printf("%v, %T\n",indexed, indexed) 
	fmt.Println("Length of the string: ", len(myString)) //returns number of bytes instead of characters, that is,   8 bytes, as é occupies 2 bytes instead of returning 6 characters
	for i,v:=range myString{
		fmt.Println(i, v)
	}

	// runes 
	//runes are an alias for int32
	var myString1=[]rune("résumé")
	var indexed1=myString1[1] //returns the actual utf-8 value , which is 233

	fmt.Println(indexed1, reflect.TypeOf(indexed1)) 

	fmt.Println("Length of the string: ", len(myString1)) //returns 6 characters
	for i,v:=range myString1{
		fmt.Println(i, v)
	}

	//Declaring a rune
	var myRune rune='a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	//Concatenating strings:
	var strSlice=[]string{"s","a","l","m","o","n"}
	var catStr=""
	for i:= range strSlice{
		catStr+=strSlice[i]
	}
	fmt.Printf("\n%v",catStr) //prints "salmon"
	//OR
	catStr=""
	for _,ele := range strSlice{
		catStr+=ele // this creates a new string every time we perform +=
	}
	fmt.Printf("\n%v",catStr) //prints "salmon"

	//Efficient way of building strings
	var strBuilder strings.Builder

	for i:=range strSlice{
		strBuilder.WriteString(strSlice[i]) //WriteString appends the contents of string(basically, strSlice[i]) to b's buffer. It returns the length of s and a nil error.
	}
	//OR
	// for _,ele:=range strSlice{
	// 	strBuilder.WriteString(ele)
	// }
	var catStr1=strBuilder.String() //returns the accumulated string(created using the buffer)
	fmt.Printf("\n%v\n", catStr1) //prints "salmon"



}