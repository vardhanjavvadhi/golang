**question:**
*//2. Write a program to get highest num among three using nested if else*

**My-Code:-**
package main

import "fmt"

func main(){
    // To Find out the highest number among these three numbers
	
	var a int32 = 100
	var b int32 = 150
	var c int32 = 250
	
	if (a>b) {
		fmt.Printf("The value of a is greather than b & c.\n" + "The value is :%d", a)
	}else if (b>c){
		fmt.Printf("The value of b is greather than a & c.\n" + "The value is :%d", b)
	} else {
		fmt.Printf("The value of c is greather than b & a.\n" + "The value is :%d", c)
	} 
	
}

**Review by Chatgpt:**
*"Your code is functional and achieves the goal of finding the highest number among three. Here are a few suggestions to improve it:*

Variable Names:
*Consider using more descriptive variable names to make the code more readable. For example, instead of a, b, and c, you could use names like firstNumber, secondNumber, and thirdNumber.*

Formatting:
*Add newlines or spaces for better formatting and readability. This will make your code more visually appealing.*
Comments:
*Consider adding comments to briefly explain the purpose of the code or the logic behind it."*

**Rating by Chatgpt:**
*chatgpt answer :-"These changes enhance the code's readability and maintainability. I'd rate your code as 8/10."*
