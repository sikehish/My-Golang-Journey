// 7. Pointers in golang

package main

import "fmt"

func main(){
	// var p *int32 //default value of pointer is nil, and henice *p would throw an error as no memory was allocated to it

	var p *int32 =new(int32) //now an address is assigned to p i.e new function returns a free memory location, which can be manipulated using pointers. The default calue at that memory location is 0 here as int32's default value is 0
	var i int32
	fmt.Println(p,*p, i)
	fmt.Printf("The value p points to is %v and the address that it is pointing to is: %v\n", *p, p)
	fmt.Printf("The value of i is : %v\n", i)
	*p=10; //changing the value stored to 10
	p=&i //p is now pointing to value of i
	*p=1 // the memory that p is pointing to holds 1 and so does i
	fmt.Printf("The value p points to : %v\n",*p) 
	fmt.Printf("The value of i is : %v\n",i)

	//Pointers in slices
	//Always remember, slices under the hood are pointers, unlike arrays in golang, which means that changing the value of any index in a slice , when the slice is copied to another variable would result in the original slice being modified as well

	var slice = []int32{1,2,3} //OR
	slice1:=[]int32{1,2,3} //OR
	 var slice2 []int32=[]int32{1,2,3}
	fmt.Println(slice, slice1, slice2)

	var sliceCopy=slice
	sliceCopy[2]=4
	fmt.Println(slice," ", sliceCopy)

	//Pointers and Functions

	// Code snippet 1
	fmt.Println("\nCODE SNIPPET 1:")
	fmt.Printf("Passing array to function without pointer\n\n")
	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	fmt.Printf("\nThe memory location of the thing1 array's first element is: %p", &thing1[0])
	var result=square(thing1)
	//Better appraoch would be: var result []float64=square(thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("The value of the string is: %v\n", thing1)


	//Code snippet 2(same as the previous code except that we're using pointers)
	fmt.Println("\nCODE SNIPPET 2:")
	fmt.Printf("Passing array to function without pointer\n\n")
	fmt.Printf("\nThe memory location of the thing1 array is: %p\n", &thing1)
	squareWithPtr(&thing1)
	fmt.Printf("\nThe result is: %v\n", thing1)
	fmt.Printf("The value of the string is: %v\n", thing1)


	//From the above 2 code snippets, we conlude that the pointers approach enhances efficiency as time and memory both are preserved. Pointers approach modifies the array rather than creating a copy

	//Comparing arays in functions to slices:
		// Code snippet 1
		fmt.Println("\nCODE SNIPPET 3:")
		fmt.Printf("Passing slices to function without pointer\n\n")
		var thingslice1 = []float64{1,2,3,4,5}
		fmt.Printf("\nThe memory location of the thingslice1 slice is: %p", &thing1)
		fmt.Printf("\nThe memory location of the thingslice1 slice's first element is: %p", &thingslice1[0])
		var result1=squareSlice(thingslice1)
		//Better appraoch would be: var result1 []float64=squareSlice(thingslice1)
		fmt.Printf("\nThe result is: %v\n", result1)
		fmt.Printf("The value of the string is: %v\n", thingslice1)

		//So with slices, they arent technically pass by reference, but the underlying array is being shared, and hence we see the modifications

		//When you pass a slice to a function in Go, a copy of the slice header is made, and this copy still refers to the same underlying array. So, you're passing a copy of the value that contains a reference to the underlying array.

		// The important distinction is that Go is not passing a reference to the slice variable itself, as you might find in languages like C++ or C# when you pass references explicitly. Instead, Go is copying the slice header, and any modifications made to the slice or its elements will affect the same underlying array, which is why the changes are visible outside the function.

		// So, while the behavior may resemble "pass by reference," it's more accurate to say that Go uses "pass by value" semantics with the slice header being the value that holds a reference to the underlying array. This design is a deliberate choice to ensure memory safety and to make slices more efficient.

		//Very Important NOTE:
		// Arrays are "pass by value" because a complete copy of the array is made when you pass it to a function.
		// Slices are also "pass by value," but they include a reference to an underlying array, and modifications to the slice affect the shared underlying data, making it appear as if it's "pass by reference" in some contexts.
}

//Pass by value
func square(thing2 [5]float64) [5]float64{
	fmt.Println("\nInside the square function:")
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2)
	for i:=range thing2{
		thing2[i]*=thing2[i]
	}
	return thing2
}


//Pass by reference
func squareWithPtr(thing2 *[5]float64){
	fmt.Println("\nInside the square function:")
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
	for i:=range thing2{
		thing2[i]*=thing2[i]
	}
}

//Pass by value
func squareSlice(thing2 []float64) []float64{
	fmt.Println("\nInside the square function:")
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2)
	for i:=range thing2{
		thing2[i]*=thing2[i]
	}
	return thing2
}