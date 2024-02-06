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
