package main

import (
	"bufio"
	"fmt"
	"os"
)
func Keyboardinput(){
	//golang can read keyboard input from the console. In this section you will learn how to do that..

//To get keyboard input, first open a console to run your program in. Then it will ask for keyboard input and display whatever you’ve typed.

//Keybpard input
reader:=bufio.NewReader(os.Stdin)
fmt.Print("Enter ypur city:")
city,_:=reader.ReadString('\n')
fmt.Print("you live in "+city)
/*
necessary
reader := bufio.NewReader(os.Stdin)

read line from console
city, _ := reader.ReadString(‘’) ``` You can get as many input variables as you want, simply by duplicating this line and changing the variable names.
*/
}