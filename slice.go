package main

import "fmt"

func Slice() {
	/*
	   The slice provides a dynamic window of an array (a slice is to an array as a slice is to a pizza) The items identified by the terminating index are not included in the slice.

	   However, unlike the array: the length of the slice can be modified at run time, with a minimum of zero maximum of the length of the array

	   The length of the slice will never exceed its capacity, so that the inequality is always true for slices: 0 <= len (s) <= cap (s)
	   Attention Never use a pointer to point to the slice, the slice itself is already a reference type, so it is itself a pointer!
	   One slice defaults to nil before it is initialized and its length is 0.
	   The initialization format of the slice is:

	   var slice1 []type = arr1[start:end]
	   Slice can also be initialized in a similar array:

	   var x = []int{2, 3, 5, 7, 11}
	   When the relevant array is not yet defined, we can use the make() function to create a slice while creating a good correlation array:

	   var slice1 []type = make([]type, len)
	   It can also be abbreviated to slice1 := make([] type, len), where len is the length of the array and also the initial length of the slice.

	   The make use mode is: func make([] T, len, cap) where cap is an optional parameter.

	   v := make([]int, 10, 50)
	   This allocates an array with a 50 int value, and creates a slice v with a length of 10 with a capacity of 50, which points to the first 10 elements of the array.

	   We have listed three types of slice initialization, which are common in all three ways.
	   If you generate a new slice from an array or slice, we can use the following expression:

	   a[low : high : max]
	*/

	//define array
	a := [5]int{1, 2, 3, 4, 5}
	//create slice from array
	t := a[1:3:5]
	//output slice
	fmt.Println(t)//[2,3]
}