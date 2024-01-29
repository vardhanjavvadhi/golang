package main

import "fmt"

func main(){
	var a complex64 = complex(8,10)
	var b complex64 = 8+9i
	var realnumber = real(a)
	var imaginary = imag(b)
	fmt.Printf("the Real number of a is : %v\n", realnumber)
	fmt.Printf("the Imaginary number of b is : %v", imaginary)

}