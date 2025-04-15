package main

import "fmt"

//varadic functions can be called with any number of arguments
//A function is said to be variadic, if the number of arguments are not explictly defined.
func Variadicfunctions(){
sum111(2,3,4)
}

func sum111(numbers ...int){
	total:=0
	for _,num:=range numbers{
		total+=num 
	}
	fmt.Println(total)
}