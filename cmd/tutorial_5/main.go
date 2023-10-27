// Strings, bytes and runes

package main

import (
	"fmt"
	"reflect"
)

//NOTE: ascii uses 7 bits while utf-8 uses 8 bits. In UTF-8, upto ascii code point 127(starting from 0), 1 byte is used to represent it in utf-8, with the MSB bit set to 0. Hence, only 0 to 127 numbered characters can be stored using utf-8 as the 1 bit is always 0
//REFERENCE: https://www.youtube.com/watch?v=8uiZC0l4Ajw Timestamp: 28:18 and 29:05

func main(){
	var myString="résumé"
	var indexed=myString[0] //returns uint8 value which is basically the utf-8 value corresponding to the character

	fmt.Println(indexed, reflect.TypeOf(indexed)) 
	//OR
	fmt.Printf("%v, %T\n",indexed, indexed) 
	fmt.Println("Length of the bytes: ", len(myString)) //returns 8 bytes, as é occupies 2 bytes
	for i,v:=range myString{
		fmt.Println(i, v)
	}

}