// Interfaces and structs

package main

import "fmt"

// type gasEngine struct{
// 	mpg uint8
// 	gallons uint8
// }
// type gasEngine struct{
// 	mpg uint8
// 	gallons uint8
// 	ownerInfo owner
// }
type gasEngine struct{
	mpg uint8
	gallons uint8
	owner
}

type owner struct{
	name string
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}


//Method
//Method on the gasEngine struct
func(e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}

func(e electricEngine) milesLeft() uint8{
	return e.kwh * e.mpkwh
}

//interfaces
type engine interface{
	milesLeft() uint8
}

// Method vs function
//Function
func milesLeft(e gasEngine) uint8{
	return e.gallons * e.mpg
}

//Older approach, only specific to gasEngine
// func canMakeIt(e gasEngine, miles uint8){
// 	if miles==e.milesLeft(){
// 		fmt.Println("We can make it!")
// 	}else{
// 		fmt.Println("We can't :((")
// 	}
// }

//Using engine interface-> can be used by both gasEngine and electricEngine as both of them have milesLeft method which returns uint8
func canMakeIt(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("We can make it!")
	}else{
		fmt.Println("We can't :((")
	}
}
func main(){

	// var myEngine gasEngine
 	// fmt.Println(myEngine.mpg, myEngine.gallons)
	// //  Will print 0 as default value of uint8 is 0)

	// var myEngine gasEngine= gasEngine{25, 15, owner{"Alex"}}
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	var myEngine gasEngine= gasEngine{25, 15, owner{"Alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	
	//Anonymous structs -> these arent usable
	var myEngine2= struct{
		mpg uint8
		gallons uint8
		}{25,15}
		fmt.Println(myEngine2.mpg, myEngine2.gallons)

		//usage of methods(and contrast with functions)
		fmt.Printf("[Method] Total miles left in tank: %v\n", myEngine.milesLeft())
		fmt.Printf("[Function] Total miles left in tank: %v\n", milesLeft(myEngine))

		var myEngine3= electricEngine{25,15}
		canMakeIt(myEngine,50)
		canMakeIt(myEngine3,50)
}