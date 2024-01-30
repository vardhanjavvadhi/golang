**Question**
*//1. Write a program to take marks as input from user and print the grade you got?*

**My-code:-**
package main

import "fmt"

func main() {
var percent float32
fmt.Printf("Enter your Percentage: ")
fmt.Scanf("%f\n",&percent)

     //To find out the grading of percent given by User

    if (percent>=90) {
       fmt.Printf("Your are passed with Distinction marks & Percentage is %f%" +"%",percent)
    } else if (percent>=65 && percent<90) {
    	fmt.Printf("Your are passed with First class marks & Percentage is %f%" + "%",percent)
    }else if (percent>=40 && percent<65) {
    	fmt.Printf("Your are passed with second class marks & Percentage isf %f%" + "%",percent)
    }else{
        fmt.Printf("Your are failed and you got only %f%" + "%",percent)
    }

}

**Review by Chatgpt:-**
Your program looks mostly correct, but there are a few issues and improvements that could be made:

Percentage Calculation:
*It seems like you want to use the percent variable as the input percentage, but the print statements have a formatting issue. Remove the extra percentage signs from the print statements.*

Spacing in Printf Statements:
*There are unnecessary spaces in your printf statements. Fix the formatting for better readability.*

Comments:
*Adding comments can make your code more understandable. You may consider adding comments to describe each section of your code.*

**Rating by chatgpt:-**
*chatgpt answer :-"I'd rate the code as 7/10."*
