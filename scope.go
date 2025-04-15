package main

import "fmt"

//define global variable
var g int

//global variable
var g1 int=20

/* Declare global variables */
var a5 int = 20;
func Scope() {

	/*
			Scope is where a variable can be used.
			A variable defined in the function is called a local variable. Their scope is only in the function body, that means they only exist within their function.

		A variable defined outside the local scope is called a global variable. Global variables can be used throughout the package or even the external package (after export).

		A variable in the function definition is called parameter
	*/

	//local variables
	//local varialbes
	var a, b, c int
	//initilize variables
	a = 10
	b = 20
	c = a + b
	fmt.Printf("variables: a=%d,b=%d and c =%d\n",a,b,c)

	//global variables
	//define local variables
	var a1,b1 int 
	//initialze variables
	a1=10
	b1=20
	g=a1+b1 
	fmt.Printf("values:a1=%d,b1=%d and g=%d\n",a1,b1,g)

	/*
	Variable priority
The global variable in the Go language program can be the same as the local variable name, but local variables within the function are prioritized.Examples are:
*/
var g1 int =10
fmt.Printf("value : g1=%d",g1)


/* main() local variables */
var a5 int = 10
var b5 int = 20
var c5 int = 0

/*
Parameters
Function Parameters (formal arguments) are used as local variables of the function. Here are some examples:
*/
fmt.Printf("main() a5 = %d\n",  a5);
c5 = sum(a5,b5);
fmt.Printf("main() c5 = %d\n",  c5);

}


/* Function definition - two numbers added */
func sum(a, b int) int {
	fmt.Printf("sum() in the function a = %d\n",  a);
	fmt.Printf("sum() in the function b = %d\n",  b);
	return a + b;
 }