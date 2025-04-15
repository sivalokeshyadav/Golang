package main

import "fmt"

func Loop() {
	/*
		syntax
		for initalization statement;conditional statement;modified statement{}
		The heads of these 3 sections,which are separated by a semicolon;they do not need to be enclosed in parantheses()
		you can use a counter in aloop, this counter i repeats until x>=10

		for i:=0;i<10;i++{}
		In the syntax above we use i++ to increase i,but i=i+1 is also permitted.Its the same thing
	*/

	for x := 0; x < 4; x++ {
		fmt.Printf("iteration x:%d\n",x)
	}

	// /Loop over array:-An array is a set of items of the same data type. You can loop over an array too:
	//define array
	a:=[]int{1,2,3,4,5}
	for i:=0;i<len(a);i++{
		fmt.Println("Character:",a[i])
	}
	for i:=0;i<len(a);i=i+1{
		fmt.Println("Character:",a[i])
	}
}