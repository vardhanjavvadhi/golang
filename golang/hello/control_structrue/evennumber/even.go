//5. Write a program to print even numbers in range 1-20

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