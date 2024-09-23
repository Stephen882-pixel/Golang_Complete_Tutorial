package main

import (
	"fmt"

	"golang.org/x/text/cases"
)
func main(){
	// Conditional statements are used to perform differen actions based on different conditions 
	// Go has the following conditional statements:
		// Use if to specify a block of code to be executed,if a condition is true 
		// Use else to specify a block of code to be executed,if a condition is false
		// use else if to specify a new condition to test, if the first condition is false 
		// Use switch to specify many alternative blocks of codes to be executed


		// if statements

		if 20 > 18{
			fmt.Println("20 is greater than 18")
		}

		sum1  := 300
		sum2  := 400
		if sum1 < sum2{
			fmt.Println("Sum2 is greater than sum1")
		}

		// if else statement

		marks := 200
		if marks > 250{
			fmt.Println("You have passed the exam")
		} else {
			fmt.Println("You have failed the exam")
		}

		// else if statements

		score := 30
		if score < 10{ 
			fmt.Println("You have failed")
		} else if score < 25 {
			fmt.Println("You have tried")
		} else {
			fmt.Println("You can proceed to the next level!!")
		}


		// Nested if >> this is where we have an if statement inside another if statement

		num1 := 30
		if num1 > 15 {
			fmt.Println("num1 is greater than 15")
			if num1 > 25 {
				fmt.Println("num1 is also greater than 25")
			}
		} else {
			fmt.Println("num1 is less than 10")
		}
		

		// Go switch statements 
		day := 4
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2 :
			fmt.Println("Tuesday")
		case 3 :
			fmt.Println("Wednesday")
		case 4 :
			fmt.Println("Thursday")
		case 5 :
			fmt.Println("Friday")
		case 6 :
			fmt.Println("Saturday")
		case 7 :
			fmt.Println("Sunday")
		}



		
		
}