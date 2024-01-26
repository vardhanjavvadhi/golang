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
