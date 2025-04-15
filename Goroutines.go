package main

import "fmt"

func Goroutines() {

	/*
		a goroutine is a function that can run concurrently
		you can see it as a lightweight thread.
		THe idea stems from concurrency:working on more than one taks simultaniously.

		To invoke a Go routine write "go" before function call.a5if you have a function f(string), call it as "go f(string)" to invoke it as goroutine.The function will then run asynchronously.a5

		Goroutines are light on memory, a program can easily have hundreds or thousands of goroutines.
	*/

	go f("go routine")
	f("function")
	fmt.Scanln()

}

func f(msg string) {
	fmt.Println(msg)
}