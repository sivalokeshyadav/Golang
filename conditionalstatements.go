package main

import "fmt"

func Ifstatement() {
	//golang can make choices with data.
	// This data (variables) are used with a condition: if statements start with a condition.
	// A condition may be (x > 3), (y < 4), (weather = rain).
	var x = 3
	if x > 2 {
		fmt.Printf("x is greater than 2\n")
	}else{
		fmt.Printf("condition is false (x > 2)\n");
	}
}


func WhileStatement(){
	//While loops are used when you are not sure how long code should be repeated. 
	// Think of a tv that should continue its function until a user presses the off button.
	// it repeat the condition until condition is true which could be forever
	i:=1
	max:=20
	// technically go doesnt have while, but
     // for can be used while in go.
     for i < max {
		fmt.Println(i)
	i += 1
	}
}


func Switch(){
	/*
	The switch statement lets you check multiple cases. You can see this as an alternative to a list of if-statements that looks like spaghetti. A switch statement provides a clean and readable way to evaluate cases.
	syntax:-
	switch var1{
		case val1:
			.....
		case val2:
			....
		default:
			....
	}
			(or)
	switch {
    case condition1:
        ...
    case condition2:
        ...
    default:
        ...
}
		(or)
	switch initialization {
    case val1:
        ...
    case val2:
        ... default:
        ...
} switch result: = calculcate (); {
    case result < 0: ...
    case result > 0: ... default:
        // 0}

Variable var1 can be any type, and val1 and val2 can be any value of the same type.
The type is not limited to a constant or integer, but must be the same type; 
The front braces {must be on the same line as the switch keyword.
You can test multiple potentially eligible values at the same time, 
using commas to split them, for example: case val1, val2, val3:. 
If you want to continue with the code for subsequent branches after you finish the code for each branch, 
you can use the fallthrough keyword to achieve the purpose.	
	*/
switch a:=1;{
case a==1:
	fmt.Println("The integer was ==1")
	fallthrough
case a==2:
	fmt.Println("The integer was ==2")
case a == 3:
	fmt.Println("The integer was == 3")
	fallthrough
case a == 4:
	fmt.Println("The integer was == 4")
case a == 5:
	fmt.Println("The integer was == 5")
	fallthrough
default:
	fmt.Println("default case")
}


}