//Golang program to calculate the reverse of the given number using the for loop

package main

import "fmt"

func main()  {
	var number int32 
	var reverse int32 
	var remainder int32 

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &number)

    //calcualte the reverse of the given number
	for number > 0 {
		remainder = number% 10
		reverse = reverse*10 + remainder
		number = number / 10
	}

	fmt.Printf("reverse of number: %d", reverse)
}