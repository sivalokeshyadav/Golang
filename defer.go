package main

import (
	"fmt"
	"os"
)

func who() {
	fmt.Println("Go")
}

func hello1(){
	fmt.Println("Hello")
}
func Defer() {
	/*
		Defer is a special statement in go.it schedules a function to be called after the current function has completed
		Go will run all statements in the function, then do the function call specified by defer after
	*/

	//who()
	//hello()
	//now first who called and then hello function called,but when we add defer before who,then first hello called then after it finished who called like
	defer who()
	hello1()

	/*
	If you want to create and write a file, the steps would normally be:

create file
write file
close file
With the defer statement you could write:

create file
(defer) close file after function completes
write file
In code that looks like this:
*/
f,_:=os.Create("hello.txt")
defer f.Close()
fmt.Fprintln(f,"hello world")

// This has some advantages:

// readability
// if run-time panic occurs, deferred functions are still run
// if the function has return statements, close will happen before that
}
