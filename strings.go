package main

//golang strings can be seen as a collection of characters. 
// A string can be of length 1 (one character), but its usually longer. 
// A string is always written in double qoutes.

//This means a string variable in golang can hold text: words, sentences, books

import (
	"fmt"
	"strings"
)

func Strings(){
	var str1="This is a string variable!"
	fmt.Print(str1)
// 	Multiple lines
// In Go programming, there are two ways to print multiple lines.

// Method 1. call display function Println x times,
// Method 2. use the newline character \n inside the string
//1.function calls
fmt.Println(str1)
fmt.Println(str1)
//2.use newline character
fmt.Printf("Hello World\n I am GO\n")
//Note: The zero value of type string is a string of zero length, that is an empty string "".

//String Index:- The contents of the string (bytes) can be obtained by standard indexing with an index written in parantheses [], and the index count from 0.
//eg str[0],str[i-1],str[len(str)-1]
//positive index starts from 0,-ve index starts from -1
//string conversion apply only for ASCII codes
	s:="GO string example"
	for k,v:=range s{
		fmt.Printf("k:%d,v:%c==%d\n",k,v,v)
	}
	//direct-use operators
	s1 := "hel" + "lo, "
s1 += "world!"
fmt.Println(s1)

//The general comparison operators (==,!=, <, <=,>=, >) implement the comparison of strings by byte comparison in memory. 
// You can get the byte length of the string by function len(), 
// for example: len(str).

//String join:To join strings, you can use the package strings which has the method .Join()
s2:=strings.Join([]string{"hello","world"},",")
fmt.Println(s2)
}