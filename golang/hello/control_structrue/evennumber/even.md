**Question:**
*//5. Write a program to print even numbers in range 1-20?*

**My-code:-**
package main

import "fmt"

func main()  {
	//for to find out the even beteen 1-20
	
    fmt.Printf("The Even Numbers are : \n")

	for i:=1; i<=20; i++{
		if (i % 2 == 0) {	
			fmt.Printf(" %d ",i)
		}		
	}
}

                                        
**Review by Chatgpt:-**
Formatting:
*Consider adding newlines for better readability, especially after the print statement. You could use fmt.Println() to separate the output from the prompt and to make the output more visually appealing.*

Comments:
*While the code is relatively simple, adding comments to briefly explain the purpose of the code or the logic within the loop can be helpful for someone else reading the code.*

**Rating by chatgpt:-**
*chatgpt answer :-"I'd rate the code as 8/10. The code is functional and achieves the goal of printing even numbers in the range 1-20."*
