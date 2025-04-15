package main

import "fmt"



func factorial(x uint) uint{
	if x==0{
		return 1
	}
	return x*factorial(x-1)
}
func Recursion() {
	/*Recursion functions are supported in GO. while recursive functions
	A function is recursive if it:
	1.Calls itself
	2.reaches the stop condition
	

	THe function below is not a recursive function

	func(x) int {
		if x == 0 {
			return 0
		}
		fmt.Println(x)
		return countdown(x-1)
	}
		*/

		x:=factorial(3)
		fmt.Println(x)
}