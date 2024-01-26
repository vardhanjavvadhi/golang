package main

import "fmt"

func main(){
	var a float32= 4.0 
	var b float32 = 7.0
	var c float32= 9.0
	var avg float32
	avg =(a+b+c)/3
	fmt.Printf("%T", a)
	fmt.Println(avg)
}