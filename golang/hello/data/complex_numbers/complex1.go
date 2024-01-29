package main

import "fmt"

func main(){
	var x complex64 = complex(9,10)
	var y complex64 = complex(8,16)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("Result is: %v + %v = %v\n", x , y , x+y)
	fmt.Printf("The type of a is :%T and " + "The type of b is : %T", x, y)
}