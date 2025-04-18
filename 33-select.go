package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string) // unbuffered, aka blocking
	c2 := make(chan string)
	// each channel receives a value after some amount of time
	go func() {
		time.Sleep(1 * time.Second) // runs concurrently with Sleep from second goroutine
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select { // select lets you wait on multiple channel operations, prints each value as it arrives
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}
}
