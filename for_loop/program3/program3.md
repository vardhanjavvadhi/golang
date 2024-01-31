*// Write a program to print the output like*
**// 1**
**// 22**
**// 333**
**// 4444**
**// 55555**

**My-code**

package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        for j := 1; j <= i; j++ {
            fmt.Print(i)
        }
        fmt.Println()
    }
}

**chat-gpt review and rating of this code?**


*Your program successfully prints the desired output pattern, and the logic is clear. Here's a brief review:*

**Rating: 9/10**

*The code is concise, and the logic is straightforward. However, for further improvement:*

*Consider adding comments to briefly explain the purpose of the code or any complex logic. While this code is relatively simple, comments can be helpful for clarity, especially in larger programs.*

*Ensure consistent indentation for better readability. Your code has consistent indentation, which is good.*