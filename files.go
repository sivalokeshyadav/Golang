package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//read line by line into memory
//all files contents is stores in lines[]
	func readLines(path string)([]string,error){
		file,err:=os.Open(path)
		if err!=nil{
			return nil,err
		}
		defer file.Close()

		var lines []string 
		scanner:=bufio.NewScanner(file)
		for scanner.Scan(){
			lines=append(lines,scanner.Text())
		}
		return lines,scanner.Err()
	}
func Files() {
	//file exists
	/*
			There is a golang function that checks if a file exists. This function returns true if the file exists.
		Why? If you try to open a file that doesn’t exist, at best it will return an empty string and at worst it will crash your program. That would lead to unexpected results and thus you want to check if the file exists.
	*/
	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Println("File exists ")
	}else{
		fmt.Println("File doesn't exist ")
	}
	//If no file root is specified, it will look for the file in the same directory as the code. 
	// If the file exists, it will return true.
	//  If not, it will return false.

	//ErrorChecking:- Sometimes you want to check if a file exists before continuing the program.
	//  This leads to clean code: first check for errors, if no errors continue.

	if _,err:=os.Stat("file-exists2.file");os.IsNotExist(err){
		fmt.Printf("File does not exist\n");
	}


	//read file
	//golang will read the file from the same directory as your program.if the file is in another directory,specify its path
	//if you want to read a file at once,you can use:

	b,err:=ioutil.ReadFile("read.go")
	//can file be opened?
	if err !=nil{
		fmt.Print(err)
	}

	//convert bytes to string
	str:=string(b)

	//shoe file data
	fmt.Println(str)


	//Line by Line
		//open file for reading
		//read line by line
	lines1,err1:=readLines("read2.go")
	if err1!=nil{
		log.Fatalf("readLines:%s",err)
	}
	//print file contents
	for i,line:=range lines1{
		fmt.Println(i,line)
	}


}


func WriteFile(){
	//golang can open a file for reading and writing (r+) or writing (w+). This is for files on the hard disk, not on the cloud.

	//write file in go. you don't need to set any flags.
	//This will overwrite the file if it already exists

	file,err:=os.Create("file.txt")
	if err!=nil{
		return

	}
	defer file.Close()

	file.WriteString("Write file in golang")
	/*
	This writes the golang string into the newly created file. 
	If you get an error, it means you don’t have the right user permissions to write a file (no write access) or that the disk is full.

Otherwise the new file has been created (file.txt). 
This file contains the string contents which you can see with any text editor.
	*/

	//Flags
	/*
If you use the w+ flag the file will be created if it doesn’t exist. 
The w+ flag makes golang overwrite the file if it already exists.

The r+ does the same, but golang then allows you to read files. 
Reading files is not allowed with the w+ flag.

If you want to append to a file (add) you can use the a+ flag. 
This will not overwrite the file, only append the end of the file.
	*/
}


func Rename(){
	/*
Renaming files is a common task in programming, and with Golang, it can be done easily. 
Golang provides a simple and efficient way to rename files within a directory, as well as move and rename them to a new directory. 
The file will be renamed to a new file. ### Rename in shell Now there are other ways to do this, for example on a Linux or Mac OS X system you can run the command

mv source.txt destination.txt	
*/

	src:="hello.txt"
	dst:="golang.txt"

	//rename file
	os.Rename(src,dst)
}