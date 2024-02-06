// Simple Go program to illustrate
// how to create a rune
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
