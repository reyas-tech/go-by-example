package main

import (
	"fmt"
	"time"
)

/*
 * Use channels to synchronize execution across goroutines
 */

func worker(done chan bool) { // function will be run in goroutine
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // send value to notify completion
}

func main() {
	done := make(chan bool, 1)
	go worker(done)     // start worker goroutine, giving it the channel to notify on
	fmt.Println("test") // this prints before worker because it is before receiver

	<-done // block until we receive a notification from the worker on the channel
	// without the above line, the program would exit before worker even started
}
