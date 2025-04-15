package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type House struct{
	noRooms string 
	price int 
	city string 
}
//gobal variable
var global int=10

func random2(min int,max int) int{
	return rand.Intn(max-min)+min
}

func Practice(){
	//Create a program with multiple string variables
	s1:="hello"
	s2:="world"
	fmt.Println(s1,s2)
	//Create a program that holds your name in a string.
	name:="Bob"
	fmt.Println(name)
	//keyboardinput
	//Make a program that lets the user input a name
	reader:=bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name:")
	name1,_:=reader.ReadString('\n')
	fmt.Print("You are Name is "+name1)
	//Get a number from the console and check if it’s between 1 and 10.
	reader1 := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a number: ")
    str2, _ := reader1.ReadString('\n')
	//remove newline
	str2=strings.TrimSpace(str2)
	//convert string varaible to int variable
	num,e:=strconv.Atoi(str2)

	if e!=nil{
		fmt.Println("conversion error:",str2)

	}
	if num>=1 && num<=10{
		fmt.Println("correct")
	}else{
		fmt.Println("num not in range")
	}

	//variables
	//Calculate the year given the date of birth and age
	//Create a program that calculates the average weight of 5 people.
	p1 := 50
    p2 := 30
    p3 := 20
    p4 := 25
    p5 := 35

    avg := (p1+p2+p3+p4+p5) / 5
	fmt.Println(avg)
	//scope
	//What’s the difference between a local and global variable?
	//local variable is defined within the flower brackers({}) and the scope variable is within the function we initalized the varaible with in the function,whereas glabal variable is defined outside of function ,its scope has globally we defined it as outside of function and used anywhere in the interface
//How can you make a global variable?

fmt.Println(global)

//array
//Create an array with the number 0 to 10
var arr1=[]int{0,1,2,3,4,5,6,7,8,9,10}
fmt.Println(arr1)
//Create an array of strings with names
var arr2=[3]string{"siva","lokesh"}
fmt.Println(arr2)

//for loop
//Can for loops exist inside for loops?
//yes we can use for loop inside for loop
//Make a program that counts from 1 to 10.
forarr:=[10]int{}
for i:=0;i<len(forarr);i++{
	fmt.Printf("%d:%d, ",forarr[i],i)
}

//range
//what is the purpose of range?
//range is used in conjunction with a data structure,it is used to iterate over the elements
//what the difference b/w the line for index,element  := range a and the line for _, element := range a ?
//index holds the position of the element in the data structure. Sometimes you may not need it, then you can use underscore.

//if statement
//Make a program that divides x by 2 if it’s greater than 0
var ifx=4
if ifx>0{
	ifx=ifx/2
}
fmt.Println(ifx)
//Find out if if-statements can be used inside if-statements.
//Yes,thats called nested-if

//while
//How does a while loop differ from a for loop?
//A for loop has a predefined amount of iterations. 
// A while loop doesn’t necessarily.

//files-exist
//Check if a file exists on your local disk
if _,err:=os.Stat("file-exists2.file");os.IsNotExist(err){
	fmt.Printf("File does not exist\n");
}
//Can you check if a file exists on an external disk?

//read file
//Think of when you’d read a file ‘line by line’ vs ‘at once’?
//If you have a very large file, line by line is better because otherwise it won’t fit into memory.
//Create a new file containing names and read it into an array
//Create a new file containing names and read it into an array

//struct
//Create a struct house with variables noRooms, price and city
var house House 
	house.city="hyd"
	house.noRooms="zero"
	house.price=2000
fmt.Printf("city:%s\n,noRooms:%s\n,price:%d\n",house.city,house.noRooms,house.price)

//write file
//Write a list of cities to a new file.
// file,err1:=os.Create("file.txt")
// if err1!=nil{
// 	return
// }
// defer file.Close()
// var a=[]string{"Rio","new York City","Milano"}
// for i:=0;i<len(a);i++{
// 	file.WriteString(a[i])
// 	file.WriteString("\n")
// }

//Map
// What is a map?
//map is a unordered  key-value pair variable,values must be unique and key is anytype
// Is a map ordered?:-no
// What can you use a map for?
//To store key value pairs


//randomnumbers
//Make a program that rolls a dice (1 to 6)
rand.Seed(time.Now().UnixNano())

randomNumb:=random2(1,6)

//Can you generate negative numbers?
randomNumb1:=random2(-7,7)
fmt.Printf("RandomNumber:%d-%d\n",randomNumb,randomNumb1)


//pointers
/*
Where are variables stored in the computer?

Memory

What is a pointer?

A pointer is a variable whose address is the direct memory address of another variable.

How can you declare a pointer?

Use the asterisk *, var var_name *var-type.
*/


//slices
//Take the string ‘hello world’ and slice it in two.
slices1:="hello world"
hello:=slices1[0:5]
world:=slices1[6:11]
fmt.Printf("%s-%s\n",hello,world)
//Can you take a slice of a slice?
//yes
fmt.Println(hello[0:2])



//methods
//Create a method that sums two numbers
var a float64 = 3
   var b float64 = 9
   var ret = sum1(a, b)
   fmt.Printf( "Value is : %.2f\n", ret )
   //Create a method that calls another method.
	//See above main() calls sum().

	//defer
	// Predict what this code does:

	 defer fmt.Println("Hello")
	 defer fmt.Println("!")
	 fmt.Println("World")


	//multiple-value-return
	//Change the return values from 2,4 to “hello”,“world”. Does it still work?
	//No, it won't work unless you also change the return types in the function signature to string.
	//Go is strongly typed — return types must match the declared types.
	//Can a combination of strings and numbers be used?
	//Yes — a Go function can return a combination of strings and numbers — but you need to declare that explicitly in the function signature.

	//variadic-functions
	//Create a variadic function that prints the names of students
stu:=[]string{"Alice","Bob","Charley"}
students(stu...)

//recursion

//When is a function recursive?
/*A function is recursive if it:

Calls itself
Reaches the stop condition*/
//Can a recursive function call non-recursive functions?
//yes ,it can


//Goroutine
/*
What is a goroutine?

A goroutine is a function that can run concurrently.

How can you turn a function into a goroutine?

Write go before the function call

What is concurrency?

Concurrency, do lots of things at once. That’s different from doing lots of things at the same time
*/

//Channels
/*
Channels

When do you need channels?

To communicate between goroutines

How can you send data into a channel?

c <- "message"
How can you read data from a channel?

msg := <- c
*/
}

func students(name ...string){
	for _,s:=range name{
		fmt.Println(s)
	}
}


func sum1(num1, num2 float64) float64 {
	return num1+num2
 }
