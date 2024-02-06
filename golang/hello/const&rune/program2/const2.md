# Question
//Write a program go language using const . 

# My-Code
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

# Review and Rating
*Review:*

Constant Naming: Good job on using constants A and B. They are appropriately named.

String Concatenation: Nicely done concatenating strings using the + operator. The use of += to append "!" to myWorld is also good.

Output Formatting: Clear and concise use of fmt.Println to display the concatenated string and the results of string comparisons.

String Comparisons: The use of string comparisons (== and >) is accurate and demonstrates understanding of string operations in Go.

*Rating:*
I would rate your program as 4.5/5. It's a well-written program that effectively demonstrates string concatenation and comparison using constants in Go.