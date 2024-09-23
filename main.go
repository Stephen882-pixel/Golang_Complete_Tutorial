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

	// Go Boolean datatypes
	var q bool = true //typed declaration with initial value 
	var w = false // untyped declaration with initial value 
	var e bool // typed declaration without initial value
	r := false // untyped declaration with initial value 


	fmt.Println("The value of q is :",q)
	fmt.Println("The value of w is:",w)
	fmt.Println("The value of e is:",e)
	fmt.Println("The value of r is:",r)

	// GO Integer data types
			// Integer datatypes are used to store wholes numbers without decimala eg 50,-3500,125000
			//The integer data types have two categories:Signed Integers(can store both positive and negative values),Unsigned Integer(can only store non-negative numbers)

	
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value %v",x,x)
	fmt.Printf("Type: %T, value %v",y,y)

	// Go float datatype
	// Float datatypes are used to store both positive and negative numbers with decimal places
	// There are two types of float datatypes:float32(contain 32 bits) and float64(contain 64 bit)
	// The default float is float64


	var k float32 = 34.89
	var l float64 =  3.4e+38
	fmt.Printf("Type %T, Value %v",k,k)
	fmt.Printf("Type %T, Value %v",l,l)

	// Go string data types
	// The string data type is used to store a sequence of characters(texts).They must be enclosed in double quotes

	var n string = "Ondeyo Stephen Omondi"
	var m string
	j := "My github username is Stephen882-pixel"
	fmt.Printf("Type %T, Value %v",n,n)
	fmt.Printf("Type %T, Value %v",m,m)
	fmt.Printf("Type %T, Value %v",j,j)
		


}

