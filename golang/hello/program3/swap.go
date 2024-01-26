package main

import "fmt"

func main(){
	var a=10
	var b=20

     //Before swapping two integers
	 fmt.Println("Before swapping two integer numbers:-")
	 fmt.Println("a = ", a)
	 fmt.Println("b = ", b)

	//After swapping two integers
	a,b=b,a
    fmt.Println("After swapping two integer numbers:-")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

}