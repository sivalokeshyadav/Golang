package main

import (
	"fmt"
	"time"
	 "github.com/bmuller/arrow"
)

func Timer() {
	/*

			A timer is a single event in the future. A timer can make the process wait a specified time. When creating the timer, you set the time to wait.

		To make a timer expire after 3 seconds, you can use time.NewTimer(3 * time.Second). If you only want to wait, use time.Sleep(3) instead.

		If you want to repeat an event, use tickers.
	*/

	t1 := time.NewTimer(3*time.Second)
	<-t1.C
	fmt.Println("TImer expired")


	/*
Stop timer
Unlike time.Sleep(), a timer can be stopped. In some scenarios you want to be able to cancel a process, like if you download a file or if you are trying to connect.

The program below allows a user to cancel the timer.
	*/

	go func(){
		<-t1.C 
		fmt.Println("TIme Expired 1")
	}()

	fmt.Scanln()
	stop:=t1.Stop()
	if stop{
		fmt.Println("Timer stopped")
	}


	//difference b/w 2 dates
	start := time.Date(1990, 2, 1, 3, 30, 0, 0, time.UTC)

    fmt.Println(start)
    fmt.Println(time.Now())
    
        // calculate years, month, days and time betwen dates
        year, month, day, hour, min, sec := diff(start, time.Now())

        fmt.Printf("difference %d years, %d months, %d days, %d hours, %d mins and %d seconds.", year, month, day, hour, min, sec)
        fmt.Printf("")

        // calculate total number of days
    duration := time.Now().Sub(start)
    fmt.Printf("difference %d days", int(duration.Hours()/24) )
}


func Ticker(){

	/*
Tickers are an essential tool for executing a block of code at regular intervals in Golang. 
Unlike timers, which are used for timeouts, tickers repeat the execution of a task every n seconds. 
By using a ticker, you can ensure that a specific task runs repeatedly, without having to write repetitive code.

Golang’s goroutines enable concurrent execution of tasks, and tickers can be used in goroutines to take advantage of this feature. 
By running a ticker inside a goroutine, you can execute a specific task repeatedly at regular intervals while allowing other tasks to run concurrently.
 This makes tickers a powerful tool for concurrent programming in Golang.
	*/
//tick
	go task10()
	time.Sleep(time.Second*5)

	//ticker
	go task11()
	time.Sleep(time.Second*5)
}

func task10(){
for range time.Tick(time.Second*1){
	fmt.Println("Tick")
	
}
}

func task11(){
	ticker:=time.NewTicker(time.Second*1)
	for range ticker.C{
		fmt.Println("Tick")
	}
	}

func Time(){
	current:=time.Now().UTC()
	fmt.Println("Date: "+current.Format("2006 Jan 02"))
	fmt.Println("Time: "+current.Format("03:04:05"))

	//strftime
	fmt.Println("Date: ", arrow.Now().CFormat("%Y-%m-%d"))
	fmt.Println("Time: ", arrow.Now().CFormat("%H:%M:%S"))



}



func diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}