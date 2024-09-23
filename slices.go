package main

import (
	"fmt"
)

func main(){

	// Slices are more similar to arrays but more powerfull and flexible
	// Like arrays,slices are also used to store multiple values of the same data type into a single variable
	// However,unlike arrays, the lenght of a slice can grow or shrink as you see fit
	// So in golang there are varios ways to create as slice:
			// using the []datatype{values} format
			// create a slice from an array
			// use the make() function 

	// create a slice using []datatype{values} format
	myslice1 :=[]int{}
	fmt.Println(myslice1)
	fmt.Println(len(myslice1)) // the len() function returns the lenght of the slice
	fmt.Println(cap(myslice1)) // the cap() function returns the number of elements the slice can grow or shrink to

	myslice2 := []string{"Go","Slices","are","Powerful"}
	fmt.Println(myslice2)
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))




	// create a slice from an array
	arr1 := [6]string{"Samuel","Frank","Stephen","Edwin"}
	myslice :=arr1[2:4]
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))


	// create a slice with the make() function
	myslice3 := make([]int,3,4)
	fmt.Println(myslice3)
	fmt.Println(len(myslice3))
	fmt.Println(cap(myslice3))

}