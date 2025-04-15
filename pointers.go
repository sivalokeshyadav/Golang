package main

import "fmt"

func Pointers() {

	/*
			A pointer is a variable whose address is the direct memory address of another variable.
			The asterisk * is used to declare a pointer. The form is:


		var var_name *var-type
		You have to write the variable type,


		// pointer to an integer
		//var ipointer *int

		//pointer to a float
		//var fpointer *float32
	*/
	//varaible
	var x int = 5
	//create pointer
	var ipointer *int

	//store the address of x in pointer variable
	ipointer = &x

	//display info
	fmt.Printf("Memory address of variable x: %x\n", &x)
	fmt.Printf("Memory addrfmt stored in ipointer variable: %x\n", ipointer)
	fmt.Printf("Contents of *ipointer variable: %d\n", *ipointer)

}