// Write a program to print following output
// 1
// 123
// 1234
// 12345

package main

import "fmt"

func main()  {
   	
	for i:=1;i<=5;i++{
		if i == 2{
			continue // Skip printing the "2nd line"
		}
		for j:=1;j<=i;j++{
			fmt.Print(j)
		}
		
		fmt.Println()
	}
}