package main

import "fmt"

/*
 * Channels: pipes that connect concurrent go routines
 * Can send values into channels from one goroutine and receive the values in another
 */

func main() {
	messages := make(chan string) // channels are typed by the values they convey

	go func() { messages <- "ping" }() // send a value into a channel using channel <- syntax

	msg := <-messages // <-channel receives value from the channel
	fmt.Println(msg)  // sends and receives block until both sender and receiver are ready
	// allows us to wait at end of program for message without using any other synchronization
}
