package main

import (
	"fmt"
	"time"
)

func Channels() {

	/*
		channels are a way for goroutines to communicate with eachother.
		A channel type is defined the keyword "chan"
		because goroutines run concurrently,they can't simply pass data from one goroutine to another.
		channels are needed

		How do channels work?
		if you type c<- "message", "message" gets send to the channel. Then msg:=<-c
		means recieve the message and store it in a variable msg.
	*/

	var c chan string=make(chan string)
	go f1(c)
	go f2(c)
	fmt.Scanln()

	/*
	Example 1: Unbuffered Channel
A channel is a built-in data structure that allows you to send and receive values between goroutines.
 Channels can be either buffered or unbuffered. 
 In an unbuffered channel, each send operation must be matched by a receive operation, making it a synchronous communication mechanism.
 
 In the below example, the worker() function simulates some work by sleeping for 2 seconds.
  Once the work is done, it sends a signal to the done channel by writing true to it. 
  The main() function creates the done channel and starts a worker goroutine by calling go worker(done). 
  It then waits for the worker to finish by receiving a signal from the done channel using the <-done syntax.
	*/


	done:=make(chan bool)

	go worker(done)// start a worker goroutine
	<-done // wait for the worker goroutine to finish and receive the signal
	fmt.Println("Main function exiting")


	/*
In contrast to unbuffered channels, buffered channels allow you to send a specified number of values without needing to wait for a receive operation.


In the below example, the producer() function sends 5 values to the channel c3.
 Since c3 is a buffered channel with capacity 2, the first 2 values are sent immediately, 
 and the remaining values are buffered until they can be received by the main() function’s loop. 
 Finally, the channel is closed to signal that all values have been sent. 
 The main() function receives the values from the channel using the range syntax, which automatically stops when the channel is closed.
	*/
c3:=make(chan int,2)// create a buffered channel with capacity 2
go producer(c3)

for i :=range c3{
	fmt.Println(i)
}



//ask chatgpt
// 	Buffered channel
// A buffered channel (named c) can be created with the line:

// var c chan string = make(chan string,3).

// The second parameter is the capacity. This will create a buffer with a capacity of 3.

// Then multiple messages can be stored using c <- message. If you want to output a channel element, you can use the line fmt.Println(<-c).
var c2 chan string = make(chan string,3)

    c2 <- "hello"
    c2 <- "world"
    c2 <- "go"

    fmt.Println(<-c2)
    fmt.Println(<-c2)
    fmt.Println(<-c2)
	//Because this channel is buffered, you can store values without a having another goroutine to save them immediately.

//Channel synchronization
/*
Channels can be used to synchronize goroutines. A channel can make a goroutine wait until its finished. The channel can then be used to notify a 2nd goroutine.

Imagine you have several goroutines. Sometimes a goroutine needs to be finished before you can start the next one (synchronous). This can be solved with channels.

*/
done1:=make(chan bool,1)
go task(done1)
<-done1

////Synchronizing goroutines

//Thus these goroutines (task1, task2) can be synchronous all the while running concurrently.
//Because everything is concurrent, you can still use the main thread for your program, at the same time.
done2:=make(chan bool,1)
go task1(done2)

if <- done2 {
	go task2()
	fmt.Scanln()
}


//Channel directions
/*
Channels (as function parameters) can have a direction. 
By default a channel can both send and receive data, like func f(c chan string).

But you can define a channel to be receive-only or send-only. 
If you then use it in the other direction, it would show a compile-time error. 
This improves the type-safety of the program.

//Receive only

//send only channel
A send only channel can set its values in a function, but cannot receive.

Change function to func f(c chan <- string). Hint: Here the only change is the position of <-
*/

//Receive only
c4 := make(chan string, 1)
c4 <- "hello"
f4(c4)

//send only channel
c5 := make(chan string, 1)
   f5(c5)
   fmt.Println(<-c5)

   //select in go
   /*
   Select waits on multiple channels. This can be combined with goroutines. select is like the switch statement, but for channels.

If multiple channels need to be finished before moving on to the next step, select is what you need.

Because it runs concurrently, it may happen that one task finishes before the other.

The select statement will make the program wait for both tasks to be completed, but that doesn’t mean both tasks are always finished in chronological order
   */

   //timeout
   /*
You can create a timeout with the select statement. To use timeouts with concurrent goroutines, you must import "time".

Then, create a channel with time.After() which as parameter takes time. The call time.After(time.Second) would fill the channel after a second.
   
//close

Channels can be closed. If a channel is closed, no values can be sent on it. A channel is open the moment you create it, e.g. c := make(chan int, 5).

You may want to close a channel, to indicate completion. That is to say, when the goroutines are completed it may be the channel has no reason to be open.

//Range
Iterate over buffered channel
The range keyword can be used on a buffered channel. Suppose you make a channel of size 5 with make(chan int, 5). Then store a few numbers into it channel <- 5, channel <- 3 etc.

You can then iterate over every item that was sent into the channel with the range keyword.

*/
   c6:=make(chan string)
   c7:=make(chan string)

   go f6(c6)
   go f7(c7)
   timeout := time.After(5 * time.Second)
   for  {
	select{
	case msg1:=<-c6:
		fmt.Println(msg1)
	case msg2:=<-c7:
		fmt.Println(msg2)
	case <-timeout:
		fmt.Println("Timed out. Goroutines finished.")
		return
	}

   }
   //fmt.Print("GOroutines finished.\n")


   //range
   channel:=make(chan int,5)
   channel<-5
   channel<-3
   channel<-9
   close(channel)

   for element:=range channel{
	fmt.Println(element)
   }
}

//channel select

func f6(c chan string){
	for {
		time.Sleep(1*time.Second)
		c<-"1"
	}
}
func f7(c chan string){
	for {
		time.Sleep(1*time.Second)
		c<-"2"
	}
}

//channel directions
func f4(c <- chan string){
	fmt.Println(<- c)
}

func f5(c chan <- string) {
	c <- "send only channel"
 }


//Channel synchronization
func task(done chan bool){
	fmt.Print("Running...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done<- true
}

//Synchronizing goroutines
func task1(done chan bool) {
    fmt.Print("Task 1 (goroutine) running...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func task2() {
    fmt.Println("Task 2 (goroutine)")
}


//buffered
func producer(c chan int){
	for i:=0;i<5;i++{
		c<-i// send the value to the channel
		
	}
	close(c)// close the channel after all values have been sent

}

func f1(c chan string) {
	c <- "f() was here"
}

func f2(c chan string) {
	msg := <-c
	fmt.Println("f2",msg)
}

//unbuffered
func worker(done chan bool) {
    fmt.Println("Worker started")
    time.Sleep(time.Second * 2) // simulate some work
    fmt.Println("Worker finished")

    done <- true // send a signal to the channel that the work is done
}