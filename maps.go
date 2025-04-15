package main

import "fmt"

func Maps() {
	/*
		A Golang map is a unordered collection of key-value pairs.
		For every key, there is a unique value.
		If you have a key, you can lookup the value in a map.

		Sometimes its called an associative array or hash table. An example of a map in Go:
		elements := make(map[string]string)

		You can map on other variable types too. Below an example of string to int mapping:
		alpha :=make(map[string]int)
		alpha["A"]=1

	*/

	elements := make(map[string]string)
	elements["O"] = "Oxygen"
	elements["Ca"] = "Calcium"
	elements["C"] = "Carbon"

	fmt.Println(elements["C"])

	//Hashmap:-The above creation of code works but it’s a bit over expressive. You can define a map as a block of data, in which there is the same key value mapping.
	alpha:=map[string]int{
		"a":1,
		"b":2,
		"c":3,
	}
	fmt.Println(alpha)

	//store information:You can use a map to store information. We change the map, into a map of strings to maps of strings to strings.
	website:=map[string]map[string]string{
		"Google":map[string]string{
			"name":"Google",
			"type":"Search",
		},"Youtube":map[string]string{
			"name":"Youtube",
			"type":"video",
		},
	
	}
	fmt.Println(website["Google"]["name"])
	fmt.Println(website["Google"]["type"])
}