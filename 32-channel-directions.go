package main

import "fmt"

/* Can specify if channel is meant to send/receive values */
func ping(pings chan<- string, msg string) { // only send channel
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // one channel for receives, one for sends
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
