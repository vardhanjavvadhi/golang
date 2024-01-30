//4. Write a program multiplication table of number.

package main 

import "fmt"

func main()  {
	//To find ourt the Multiplication table any Number

    var number int32
	fmt.Printf("Enter a Number:")
	fmt.Scanf("%d",&number)
	
	for i:=1; i<=15; i++{
		result:= number*int32(i)
		fmt.Printf("%d * %d := %d\n",number,i,result)
	}
	
}