package main

import "fmt"

func main(){
	var x int8 = 10
	var y int8 = 8

	//Addition
	fmt.Printf("Addition: %d + %d = %d\n", x , y , x+y)

	//Subtraction
	fmt.Printf("Subtraction: %d - %d = %d\n", x , y , x-y)

	//Multiplication
	fmt.Printf("Multiplication: %d * %d = %d\n", x , y , x*y)
	
	//Division
	fmt.Printf("Division: %d / %d = %d\n", x , y , x/y)

    //Remainder
	fmt.Printf("Remainder: %d %% %d = %d\n", x , y , x%y)
}