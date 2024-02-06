# QUESTION:-

//Golang program to create a simple calculator using switch case
**My-Code**
package main

import "fmt"

func main(){
var number1 float64
var number2 float64
var choice float64
var result float64

fmt.Print("Enter the First value: ")
fmt.Scanf("%f",&number1)

fmt.Print("Enter the Second value:")
fmt.Scanf("\n%f",&number2)

fmt.Println("for Addition Press 1")

fmt.Println("for Multiplication Press 2")

fmt.Println("for Subtraction Press 3")

fmt.Println("for Division Press 4")

fmt.Println("for Remainder Press 5")

fmt.Println("Enter your choice: ")
fmt.Scanf("\n%f",&choice)

switch choice {
case 1:
result =number1+number2
fmt.Println("Then the result is : ", result)

case 2:
result =number1*number2
fmt.Println("Then the result is : ", result)

case 3:
result =number1-number2
fmt.Println("Then the result is : ", result)

case 4:
result =number1/number2
fmt.Println("Then the result is : ", result)

default:
   fmt.Println("Invaild Operation/value")
}

}

*Chatgpt review and rating:-*
Your code appears to be a simple calculator program in Go. Here's a review:

Readability and Structure:

The code is well-structured and easy to read.
Variable names are clear, and comments are not necessary due to the simplicity of the program.
User Input:

It correctly takes two floating-point numbers as input.
Menu and Choice Handling:

The menu for the user to choose the operation is clear and well-presented.
The switch statement effectively handles different choices, performing the corresponding operations.
Output:

The output statements in each case are clear and include the calculated result.
Improvement Suggestions:

You might want to handle the case where the user attempts division by zero. Adding a check for this situation could enhance the robustness of your program.
Here's an updated version of your code with the suggested improvement:

Rating: 4.5 out of 5.
The code is well-written and achieves its purpose effectively. The suggested improvement adds a check for division by zero, which is a common edge case in calculator programs.
