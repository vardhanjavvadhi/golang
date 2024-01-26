package main

import "fmt"

func main(){
	var a int = 10
	fmt.Println("Enter an integer: ")
	fmt.Scanln(&a)
	fmt.Println("a = ", a)
}