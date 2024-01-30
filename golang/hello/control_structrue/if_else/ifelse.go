//1. Write a program to take marks as input from user and print the grade you got.

package main

import "fmt"

func main()  { 
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
