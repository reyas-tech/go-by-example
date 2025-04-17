package main

import (
	"fmt"
	"time"
)

/* Goroutine: lightweight thread of execution, allows for asynchronous execution */

func f3(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f3("direct") // run synchronously

	go f3("goroutine") // invoke f3 in goroutine
	// new goroutine runs concurrently with previous call

	go func(msg string) { // can start goroutine for anonymous function call
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // wait for the 2 goroutines to finish
	fmt.Println("done")
}
