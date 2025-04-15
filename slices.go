package main

import "fmt"

func Slices() {
	//A slice can be a subset of an array, list or string. You can take multiple slices out of a string, each as a new variable.
	//list/set of numbers
	myset := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	//take slice
	s := myset[0:4]
	fmt.Println("slcies:",s)

	//string slice
	myString:="Go programming"
	s1:=myString[0:2]
	fmt.Println("slcies:",s1)


}