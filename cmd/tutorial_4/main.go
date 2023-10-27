package main

import "fmt"

//4 points about Arrays in Golang:
// 1) Fixed length
// 2) Stores elements of same type
// 3) Indexable
// 4) Contiguous in memory
// 5) Default values in the array depends on the default value of the element types on which the array is based. In this case, for int32, it is 0

func main(){

	// Arrays
	var intArr [3]int32
	intArr[1]=132
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3]) //access 1st and 2nd element 

	fmt.Println(&intArr[0])

	//OR var intArr [3]int32=[3]int32{1,2,3} OR intArr:= [3]int32{1,2,3} OR intArr:=[...]int32{1,2,3}

	//Slices
	//Slices wrap arrays and provide extra functioanlity(it isnt fixed size and hence modifiable)
	var intSlice []int32=[]int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice=append(intSlice,7) //after appending, ta new arr is returned(also if capacity==length, then capacity is doubled/changed, but the slots unoccupied in the array still cant be accessed despite the capacity being changed)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	// fmt.Println(intSlice[4]) -->will throw an error

	var intSlice2 []int32=[]int32{8,9}
	intSlice=append(intSlice,intSlice2...) //spread operator applied on intSlice2
	fmt.Println(intSlice)


	//Slices can be created using make function
	var intSlice2 []int32=make([]int32, 3, 100) //3 is length and 100 is capacity

	fmt.Printf("The length is %v with capacity %v", len(intSlice2), cap(intSlice2))
	//Accessing beyond the capacity will not give any errors but it may lead to unexpected results
	intSlice2 = append(intSlice2, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice2), cap(intSlice2))
	fmt.Println(intSlice2[4])


	//Array vs Slice
	//1) Array is value type whereas slice is reference type
	//2) Array is fixed length whereas slice is dynamic
	//3) Array cant be resized whereas slice can be resized using append function

	//Maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	
	var myMap2 map[string]uint8 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"]) //will return 0 as default value of uint8 is 0. Always remember, that accessing keys that arent present will also yield a value, and that is always the default value

	delete(myMap2, "Adam") //deletes by ref and hence doesnt return a new map

	//checking if age exists
	var age, ok=myMap2["Jason"]
	if ok{
		fmt.Printf("The age is %v\n", age)
	}else{
		fmt.Println("Invalid Name")
	}

	//Different types of iterating
	for name:= range myMap2{
		fmt.Printf("Name: %v\n",name)
	}

	for name, age:= range myMap2{
		fmt.Printf("Name: %v Age: %v\n",name, age)
	}

	for i,v := range intArr{
		fmt.Printf("Index: %v, Value: %v \n",i,v)
	}

	var i int=0
	for i<10{
		fmt.Println(i)
		i+=1
	}

	// //OR
	// for true{
	// 	if i==10{
	// 		break
	// 	} 
	// 	fmt.Println(i)
	// 	i+=1
	// }

	// //OR
	// for i=0; i<10; i++{
	// 	fmt.Println(i)
	// }

	//NOTE: Setting capacity ahead of time (preallocation) of a slice is a lot faster(performance/time wise) than letting the capacity update dynamically
	//Reference: https://www.youtube.com/watch?v=8uiZC0l4Ajw Timestamp: 26:02

}