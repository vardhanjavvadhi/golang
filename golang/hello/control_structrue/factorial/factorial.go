//3. Write a program to find factorial of number using for loop

package main

import "fmt"

func main(){
   // to find out Factorial number
   
  var  numberToFindFactorial int64 = 15

    result := int64(1)
    for i :=int64(2); i<=numberToFindFactorial; i++ {
        result *= i
    }
    fmt.Printf("The factorial of %d is: %d\n", numberToFindFactorial, result)
}