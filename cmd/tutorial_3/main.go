package main

import (
	"errors"
	"fmt"
)

func main(){
	var printValue string="Hello World"
	printMe(printValue)

	var numerator=11
	var denominator=2

	var result, remainder, err= intDivision(numerator, denominator)

	// if err!=nil{
	// 	fmt.Printf(err.Error())
	// }else if remainder==0{
	// 	fmt.Printf("The result of the integer division is %v", result)
	// }else{
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

	//The above if else block is rewritten using switch case(also we dont use breaks in switch case in golang)

	//Switch without a variable
	switch{
	case err!=nil:
			fmt.Println(err.Error())
	case  remainder==0:
			fmt.Printf("The result of the integer division is %v\n", result)
	default:
			fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}

	//Switch with a variable(the usual case)
	switch remainder{
		case 0:
			fmt.Println("The division was exact")
		case 1:
			fmt.Println("The division was close")
		default:
			fmt.Println("The division was not close")
	}
	
}

func printMe(value string){
	fmt.Println("The value is: ", value)
}

func intDivision(nr int, dr int) (int, int , error) {
	var err error //default value is nil

	if dr==0{
		err=errors.New("Cannot divide by 0")
		return 0,0, err
	}

	var result=nr/dr
	var remainder=nr%dr

	return result, remainder, err
}

//runes are an alias for int32