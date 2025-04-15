package main

import "os"

func Panic() {
	/*
			Go has a Panic(msg) function. If the panic function is called, execution of the program is stopped. The panic function has a parameter: a message to show.

		You can use the panic function if a failure occurs that you don’t want or don’t know how to deal with.
		Panic function
		Real life scenario: Your program needs to create a file, but you don’t want to deal with error processing.

		The panic() function will make the program exit if it cannot create the file.
	*/
	_, err := os.Create("/root/example")
	if err!=nil{
		panic("cannot create file")
	}

}