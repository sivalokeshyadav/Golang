package main

import "fmt"

func Methods() {
	/*
			golang methods are reusable code blocks.
			By calling a golang method, all of the code in the method will be executed.

		Methods should start with a lowercase character and only contain alphabetic characters.
		 A method can take one or more parameters which can be used in the code block.
	*/

	hello("go")
	hello("")


	var a float64=3
	var b float64=9

	var ret =multiply(a,b)
	fmt.Printf("Value is : %.2f",ret)
}

func hello(x1 string) {
	fmt.Printf("Hello %s\n",x1)
}


func multiply(num1,num2 float64) float64{
	var result float64 
	result=num1*num2
	return result
}