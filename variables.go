package main

import (
	"fmt"
)

func Variables(){
	//Variables often hold text or numeric data. 
	// In golang there are several types of variables, 
	// including strings and numeric variables.


	//varaible Declaration 
	//Variables declaration in the Go language: 
	// variable names consist of letters, numbers, underscores, 
	// where the first letter cannot be numeric.
	/* var (
  a int 
  b bool 
  s string
  apple int
  bee bool
  apple2 string
)
This type of factoring keyword is generally 
written to declare global variables and is generally defined outside of func.

When a variable is declared by a var, the system automatically gives it a zero value of that type:
int i //= 0
float f //= 0.0
bool b //= False
string s //= " "
pointer p // = nil
and these variables are initialized in Go.

Multiple variables can be assigned on the same line, also known as parallel assignments. For example:

a, b, c = 5, 7, "abc"
// simple declaration:
a, b, c := 5, 7, "abc" 

The values on the right of the colon before the equal sign are assigned to the left variable in the same order, so the value of a is 5, the value of b is 7, and the value of c is “abc”.
*/

x:=1
fmt.Println(x)//prints 1
{
	fmt.Println(x)//prints 1
	x:=2
	fmt.Println(x)//prints 2
}
fmt.Println(x)//prints 1

//Strings:-a string is a text variable,this can be a single character, a word,a phrase, a paragraph or even a book
//It is defined in double quotes
//to print multiple values we use %s in strings
s:="Hello"
y:="World"
fmt.Printf("%s ,%s",s,y)

//Numeric Varaibles:-there are 2 main types of numeric varaibles in programming integers(int),floating points(float)
//Integer:-An integer is a natural number like 1,2,3,4. If you define a varaible to be an integer,it can't be a real number(1.5,1.3333)
a:=1
b:=2
fmt.Printf("a is %d\n && b is %d\n",a,b)
/*
%b  base 2
%c  the character represented by the corresponding Unicode code point
%d  base 10
%o  base 8
%O  base 8 with 0o prefix
%q  a single-quoted character literal safely escaped with Go syntax.
%x  base 16, with lower-case letters for a-f
%X  base 16, with upper-case letters for A-F
%U  Unicode format: U+1234; same as "U+%04X"
*/

//Float:- it is not a natural number.You can have values Pi,1.3333,1,23456......
apple:=3.0
bread:=2.0
price:=apple+bread
fmt.Printf("")
fmt.Printf("price: %f\n",price)
vat:=price*0.15
fmt.Printf("VAT:%f\n",vat)
total:=vat+price
fmt.Printf("Total: %f\n",total)
fmt.Printf("")
//we can limit numbers after dot(.) in float
fmt.Printf("Price:    %.2f\n",price)
fmt.Printf("VAT:      %.2f\n",vat)
fmt.Printf("Total:    %.2f\n",total)

//Exchange varaibles:exchanging values of 2 varabiles
//i,j=j,i 
i:=1
j:=2
i,j=j,i
fmt.Println(i)
fmt.Println(j)

//black identifier
/*The blank identifier is used to abandon the value, 
such as the value of 5 is abandoned in ` , b = 5, 7.

_, b = 5, 7
_ is actually a write-only variable and you cannot get its value. 
This is done because you must use all of the declared variables in the Go language, 
but sometimes you do not need to use all the return values you get from a function.

Because the Go language has a mandatory requirement, 
you must use the declared variable within the function,
 but the global variable that is not used is no problem.
 In order to avoid unused variables, the code will fail to compile, 
 and we can change the unused variable to _.
 */

 //nil value:-

 /*
 The nil designator is used to represent a “zero value” of interface, 
 function, maps, slices, and channels.
 If you do not specify the type of variable, 
 the compiler will not be able to compile your code because it cannot guess a specific type.
 
 var x1 = nil
 _ = x
 adding an element in a slice of nil is no problem, but doing the same thing to 
 one map will generate a runtime panic:
 var m map[string]int
m["one"] = 1 //error

string is not nil

This is a place that requires attention for developers who often use nil assignment string variables.

var str string = " " 
" " is the zero value of the string and the writing below is the same as the above:

var str string
*/

}