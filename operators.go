package main

import (
	"fmt"
)

func main(){
	// Operators are used to perfom operations of variables and values
	var a = 10 + 10
	fmt.Printf("a is of type %T, and its value is %v:",a,a)

	sum1 := 100 + 100
	sum2 := sum1 + 100
	sum3 := sum1 + sum2
	fmt.Println("The value of sum3 is:",sum3)

	// Golang has the following categories of operators
		// Arithmetic operators >> +,-,/,*,%,++,--
		// Assignment operators >> 
		// Comparison Operators 
		// Bitwise operators 
		// Logical operators
	
		sum4 := 300
		sum5 := 400
		fmt.Println(sum4>sum5)
		fmt.Println(sum4==sum5)

		fmt.Println(sum1<sum5 && sum5<sum1)


}