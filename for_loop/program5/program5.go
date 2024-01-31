//Golang program to check the given number is Prime or not using the for loop

package main 

import "fmt"

func main(){
	var num int
    var rem int

	fmt.Print("Enter Number: ")
	fmt.Scanf("%d", &num)
        //calcualte the prime number
		
	for i:= 2; i<=num/2; i++ {
		if num % i == 0 {
			rem = 1
			break
		}
	}
	if rem == 1 {
		fmt.Print("Given number is not PRIME number")
	} else {
		fmt.Print("Given number is PRIME number")
	}
}

