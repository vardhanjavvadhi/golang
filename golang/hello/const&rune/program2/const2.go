package main

import "fmt"

func main(){
	const A = "VaJ"
	const B = "VARDHANJAVADHI"
	
	// Concat strings.
	var myWorld = A+ " " + B
	myWorld += "!"
	fmt.Println(myWorld) 
	
	// Compare strings.
	fmt.Println(A == "VaJ") 
	fmt.Println(B > A) 
}
