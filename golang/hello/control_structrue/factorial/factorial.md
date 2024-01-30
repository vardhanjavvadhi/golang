**Question:**
//3. Write a program to find factorial of number using for loop

**My-Code**
package main

import "fmt"

func main(){
   // to find out Factorial number
   
  var  numberToFindFactorial int64 = 15

    result := int64(1)
    for i :=int64(2); i<=numberToFindFactorial; i++ {
        result *= i
    }
    fmt.Printf("The factorial of %d is: %d\n", numberToFindFactorial, result)
}

**Review by Chatgpt:-**
Positive Aspects:

Correctness: 
*The code correctly calculates the factorial of a given number using a for loop.*

Efficiency: 
*The use of a for loop to calculate the factorial is efficient and appropriate.*

Suggestions for Improvement:

Variable Naming: 
*While the variable name numberToFindFactorial is descriptive, consider shorter names like n for brevity and clarity.*

User Input: 
*If you want to make the code more interactive, you can allow the user to input the number for which they want to find the factorial. This can be done using fmt.Scan to read user input.*

**Rating by Chatgpt:-**
*chatgpt answer :-"I'd rate your code as 9/10."*