package main

import (
	"fmt"
)

func main(){
	var a bool = true
	var b int = 5
	var c float32 = 3.14
	var d string = "Stephen"

	fmt.Println("Boolean value is:",a)
	fmt.Println("The integer value is:", b)
	fmt.Println("The float32 value is:",c)
	fmt.Println("The string value is:",d)


	var q bool = true //typed declaration with initial value 
	var w = false // untyped declaration with initial value 
	var e bool // typed declaration without initial value
	r := false // untyped declaration with initial value 


	fmt.Println("The value of q is :",q)
	fmt.Println("The value of w is:",w)
	fmt.Println("The value of e is:",e)
	fmt.Println("The value of r is:",r)
	
}

