# question:-

//Golang program to demonstrate the switch case without optional statement and expression.

package main

import "fmt"

func main(){
var day int32 =7
switch day {
case 1:
fmt.Println("Monday")

case 2:
fmt.Println("Tuesday")

case 3:
fmt.Println("Wednesday")

case 4:
fmt.Println("Thursday")

case 5:
fmt.Println("Friday")

case 6:
fmt.Println("Saturday")

case 7:
fmt.Println("sunday")

default:  
fmt.Println("Invalid day")
}

}

*ChatGpt review and rating:*

Your Go program demonstrates the use of a switch statement without an optional statement or expression. Here's a review:

Readability and Structure:

The code is straightforward and easy to read.
Variable names are clear, and the switch statement is well-organized.
Correctness:

The switch statement correctly maps the value of day to the corresponding day of the week.
The default case handles the situation where the value of day doesn't match any of the defined cases.
Consistency:

You've used consistent indentation and formatting, which is good for readability.
Improvement Suggestions:

No significant improvements are needed for the given program. It achieves its purpose effectively.
Here's a revised version with minor formatting adjustments:

Rating: 4.5 out of 5.
The code is well-written, and it effectively demonstrates the use of a switch statement for day-of-week mapping. It follows best practices and is easy to understand.
