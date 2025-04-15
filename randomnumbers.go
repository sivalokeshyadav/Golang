package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


func random (min int,max int) int{
	return rand.Intn(max-min)+min
}
func RandomNumbers() {
	/*
		In Golang, the math/rand package provides functions for generating random numbers.
		Here’s a quick guide on how to use this package to generate random numbers.

		Setting the Seed
		Before generating random numbers, you need to set the seed value.
		The seed value is used as the starting point for generating random numbers.
		If you set the same seed value, you’ll get the same sequence of random numbers every time you run your program.
		To set the seed, use the

		rand.Seed() function:

Generating Integers
To generate a random integer, use the rand.Intn() function.
 This function takes a single argument, which is the maximum value for the random number (exclusive).
	*/

	rand.Seed(time.Now().UnixNano())
	//generating random Numbers here
	randomInt:=rand.Intn(10)
	
	fmt.Println(randomInt)

	//float numbers
	randomFloat:=rand.Float64()
	fmt.Println(randomFloat)

	/*
	Generating Unique IDs
One common use case for random numbers is generating unique IDs.
 To generate a unique ID, you can use a combination of random integers 
 and the strconv package to convert the integers to strings
*/

id:=""
for i:=0;i<8;i++{
id+=strconv.Itoa(rand.Intn(10))
}
fmt.Println(id)


//random number in golang
randomNum:=random(0,10)
randomNum1:=random(20,40)
fmt.Printf("Random Number:%d-%d\n",randomNum,randomNum1)
}