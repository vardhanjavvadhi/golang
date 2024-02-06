# Question:
//Write a program on rune to Print UTF-8 ASCII.

# My-Code:
package main

import "fmt"

func main() {
   // str := "My name is vardhan"
   str := "My"
 
     // Print ASCII representation
    fmt.Printf("ASCII representation: %x\n", str)
    // Print UTF-8 representation with the index
    for i, char := range str {
        fmt.Printf("Index: %d, UTF-8 representation for %c: %U\n", i, char, char)
    }
}

# Review and Rating given by Chatgpt: 
*Review:*

Code Structure: Your code structure is clear and concise, making it easy to understand.

Variable Naming: Good job on using meaningful names like str, i, and char.

Comments: It's good practice to include comments. However, in this case, the comment mentioning "Print UTF-8 ASCII" might be a bit misleading. UTF-8 and ASCII are different character encodings. You are printing the ASCII representation and UTF-8 representation of the string.

Unused Code: You have commented out a line of code (// str := "My name is vardhan") which is fine, but it's always good to remove unnecessary or commented-out code in the final version.

*Rating:*
I would rate your program as 4/5. It effectively demonstrates printing ASCII and UTF-8 representations of a string, and the code is well-structured.
