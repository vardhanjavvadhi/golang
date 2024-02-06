# Question:
// Simple Go program to illustrate how to create a rune

# My-Code
package main

import "fmt"

func main() {

	// Creating a rune
	rune1 := 'B'
	rune2 := 'g'
	rune3 := '\a'

	// Displaying rune and its type
	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %T", rune1,rune1,rune1)
	
	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %T", rune2,rune2, rune2)
	
	fmt.Printf("\nRune 3: %c; Unicode: %U; Type: %T", rune3, rune3, rune3)

}

# Review and Rating given by Chatgpt:


*Review:*

Variable Naming: Good job on choosing clear and meaningful names for your rune variables (rune1, rune2, and rune3).

Rune Representation: You correctly use %c to print the character, %U to print the Unicode code point, and %T to print the type.

Unicode Escape Sequence: The use of \a as a Unicode escape sequence is a nice touch to show a non-printable character.

*Rating:*
I would rate your program as 4/5. It effectively demonstrates how to create runes in Go and displays their characters, Unicode code points, and types.