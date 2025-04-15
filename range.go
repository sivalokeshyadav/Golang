package main

import "fmt"

func Range() {
	/*
		Range is always used in conjunction with a data structure.
		Thus, the first step is to create a data structure.
		nums:=[]int {1,2,3,4,5,6}

	*/
	nums := []int{1, 2, 3, 4, 5, 6}
	for _, num := range nums {
		fmt.Println(num)
	}

	//Index
	var a=[]int64{1,2,3,4}
	for index,element:=range a{
		fmt.Print(index,")",element,"\n")
	}

}