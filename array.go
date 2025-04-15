package main

import "fmt"

func Array() {
	/* An array is a set of numbered and length-fixed data item  sequences that have the same unique type
	This type can be any primitive type such as : integer,string,custom type
	The array length must be a constant expression and must be a non-negative integer.
	An array in the Go language is a type of value (not a pointer to the first element in c/c++), so it can be created by new()
	The array element can be read (or modified) by an index (position), the index starts from 0, the first element index is 0, the second index is 1, and so on.
	the array length is upto 2 Gb
	var arr1=new([5]int)
	fmt.Println(arr1[0])//output first element
	The method of traversing an array can either be a condition loop or a for-range can be used.Both of these structures are equally applicable to Slices.
	*/
	var arr1 = new([5]int)
	fmt.Println(arr1[2])
	fmt.Println(arr1[3])
	//var arrAge=[5]int{18,20,15,22,16}
	//var arrlazy=[...]int{5,6,7,8,22}
	//var arrKeyValue=[5]string{3:"Chris",4:"Ron"}

	var a=[]int{1,2,3,4,5}
	fmt.Printf("First element %d\n",a[0])
	fmt.Printf("Second element %d\n",a[1])
	//The index should not be larger than the array, that could throw an error or unexpected results.



}