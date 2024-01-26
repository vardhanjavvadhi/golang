package main

import "fmt"

func main() {
	 // Demonstrate Printf() for formatted output
name := "vardhan"
age := 25

fmt.Printf( "Hello, %s! You are %d years old.\n", name,age)

    // Demonstrate Scanf() for formatted input
	var inputName string
	var inputage int

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s\n",&inputName)

	fmt.Print("Enter your age: ")
	fmt.Scanf("%d\n",&inputage)

	

	fmt.Printf("Nice to meet you, %s! You are %d years old.\n ", inputName, inputage)
}
//Golang program to demonstrate the use of 
//printf() and scanf() function.

// package main
// import "fmt"

// func main() {
//     //Declare 4 integer type variables.
//     var p int
//     var r int
//     var t int
//     var si int
    
    
//     fmt.Print("Enter principal: " ) 
//     fmt.Scanf("%d\n",&p)
    
//     fmt.Print("Enter rate: ") 
//     fmt.Scanf("%d\n",&r)
    
//     fmt.Print("Enter time: ") 
//     fmt.Scanf("%d\n",&t)
    
//     si=(p*r*t)/100
    
//     fmt.Println("Simple interest is: ",si) 
// }
