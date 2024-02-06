//Golang program to create a simple calculator using switch case
package main

import "fmt"

func main(){
	var number1 float64
	var number2 float64
	var choice float64
    var result float64

  fmt.Print("Enter the First value: ") 
  fmt.Scanf("%f",&number1)

  fmt.Print("Enter the Second value:") 
  fmt.Scanf("\n%f",&number2)

  fmt.Println("for Addition Press 1")	

  fmt.Println("for Multiplication Press 2")	

  fmt.Println("for Subtraction Press 3")	

  fmt.Println("for Division Press 4")	

  fmt.Println("for Remainder Press 5")	


  fmt.Println("Enter your choice: ")	
  fmt.Scanf("\n%f",&choice)

  switch choice {
  case 1:
	result =number1+number2
	fmt.Println("Then the result is : ", result)
	
case 2:
	result =number1*number2
	fmt.Println("Then the result is : ", result)

case 3:
	result =number1-number2
	fmt.Println("Then the result is : ", result)

case 4:
	result =number1/number2
	fmt.Println("Then the result is : ", result)


default:
   fmt.Println("Invaild Operation/value")	
  }


}
