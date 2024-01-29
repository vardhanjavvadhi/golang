package main 

import "fmt"

func main(){
	//for 8bit unsinged integer
	var X uint8 =22
	fmt.Println("UnSigned integer :- ", X , 1-X , X+1)
	//for 8bit singed integer
	var Y int8 =48
	fmt.Println("Signed integer :- ", Y , 2-Y ,Y+3)

	 // Using 16-bit signed int 
	 var Z int16 = 32767
	 fmt.Println(Z-2, Z+2) 
}