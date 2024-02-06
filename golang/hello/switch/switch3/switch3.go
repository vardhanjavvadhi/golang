// Go program to illustrate the concept of Expression switch statement 

package main 

import "fmt"

func main() { 
	var value string = "five"
	
	// Switch statement without default statement 
	// Multiple values in case statement 
switch value { 
	case "one": 
	fmt.Println("C") 

	case "two": 
	fmt.Println("Javascrpit") 
	
	case "three","four", "five": 
	fmt.Println("Java") 
    
    case "six", "seven":
     fmt.Println("Python")

		
} 
} 
