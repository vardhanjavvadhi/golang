package main

import "fmt"

func main(){
	var string1 string ="hello world" 
	var string2 string ="Myself Vardhan_Javvadhi"
	var string3 string ="hello world"

	 result1:= string1 == string2 
	 result2:= string1 == string3
	fmt.Printf("the result1 is %t\n",result1)
    fmt.Printf("The result2 is %t\n",result2)

	fmt.Printf("The type of result1 is %T and "+"The type of result2 is %T",result1,result2)
}