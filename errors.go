package main

import (
	"errors"
	"fmt"
)

func Errors() {
	/*
		Go has a builtin type for errors. In Go, an error is the last return value (They have type error).

		The line errors.New creates a new error with a message.
		 If there is no error, you can return the nil value.


	*/

	fmt.Println(do())
	r, e := do()
if r == -1 {
    fmt.Println(e)
} else {
    fmt.Print("Everything went fine\n")
}
/*
This example demonstrates how to use the do() function to return both the return value and an error message. 
The r variable captures the return value, while the e variable captures the error message. 
The if statement then checks if the return value is equal to -1. 
If it is, the error message is printed out. Otherwise, the message “Everything went fine” is printed out.
*/
}

func do() (int, error) {
    return -1, errors.New("something wrong")
}