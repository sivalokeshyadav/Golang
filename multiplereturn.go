package main

import "fmt"

func values() (int, int) {
	return 2, 4
}
func Multiplereturn() {
	/*
		In Go, a function can return one or more values.
		This is a powerful feature not found in all programming languages.
		 Traditionally, many programming languages only allow functions to return zero or at most one value.
		  But with Go, you can return as many values as you need, making it easier to work with complex systems.


		  When you call a function that returns a value, the value is typically stored in a variable within the function’s scope.
		  However, when you return that value, it can be used outside the function.
		  This is useful for accessing variables that might not be visible in the rest of the program.
		  By returning them from a function, they can be used and manipulated more easily.

		Overall, returning multiple values in Go is a powerful feature that allows developers to write more effective and concise code.
		 In this article, we will explore how to write functions that return multiple values, and how to use them in practice.
	*/

	values()
	x, y := values()
	fmt.Println(x)
	fmt.Println(y)
}