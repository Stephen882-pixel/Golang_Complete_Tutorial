package main

import (
	"fmt"
)

func main(){
	// Arrays are user to store multiple values of the same data type into a single variable instead of having to declare multiple variables 

	var set1 = [5]int{123,123,123,123,123}
	set2 := [4]string{"Ondeyo","Stephen","Omondi","Stephen882"}
	set3 := [...]int{123,123,123,123,123,123} //the lenght of this array is infered

	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(set3)

	// Accessing elements of an array
	fmt.Println(set2[1])

	// Changing elements in an array

	set2[0] = "Stefan881"
	fmt.Println(set2)


	// Initialize only specific elements 
	set4 := [5]string{1:"Newton",4:"Odhiambo"}
	set5 := [5]int{1:10,4:30}
	fmt.Println(set4)
	fmt.Println(set5)


	// find the length of an array
	fmt.Println("The lenght of set2 is:",len(set2))



}
