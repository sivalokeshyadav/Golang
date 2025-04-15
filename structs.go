package main

import "fmt"

type Person struct {
	name string
	job  string
}

func Structs() {
	/*
		A struct can bundle attributes together. If you create a struct, you can set a couple of variables for that struct. Those variables can be of any datatype.

		A key difference from an array is that elements of an array are all of the same datatype. That is not the case with a struct.

		If you want to combine variables, structs are the way to go. Unlike the concept of object oriented programming, they are just data holders.
	*/
	var aperson Person
	aperson.name = "Albert"
	aperson.job = "Professor"

	fmt.Printf("aperson.name=%s\n",aperson.name)
	fmt.Printf("aperson.job=%s\n",aperson.job)

}