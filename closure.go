package main

import "fmt"

func Closure() {
	/*
		closure :- A closure is a type of function, that uses variables defined outside of the function itself.

	*/
	world := func() string {
		return "world"
	}

	fmt.Print("hello",world(),"\n")


	x:=2
	increment:=func ()int{
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

	//Generator
	//You can use the idea of closures to make a number generator.
//In the example below function makeSequence() returns an anonymous function that generates odd numbers.
//  An anonymous function is just a function without a name.

sequenceGenerator := makeSequence()
fmt.Println(sequenceGenerator())
fmt.Println(sequenceGenerator())
}


func makeSequence() func() int {
    i:=1
    return func() int {
        i+=2
        return i
    }
}