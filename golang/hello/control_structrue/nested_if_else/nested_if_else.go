//2. Write a program to get highest num among three using nested if else

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
