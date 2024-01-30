package main

import "fmt"

func main(){
	var age int32
	fmt.Printf("Enter your age: ")
	fmt.Scanf("%d",&age)
	
if (age>18){
	fmt.Printf("The canditate should be adult")
}else if(age>12 && age<18) {
	fmt.Printf("The canditate should be teenger")
}else
{
    fmt.Println("The canditate should kid")
}
}