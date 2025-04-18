package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages: // non-blocking receive
		fmt.Println("received message", msg)
	default: // will immediately take default case if no value on <-messages case
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg: // non-blocking send
		fmt.Println("sent message", msg)
	default: // default selected because msg can't be sent because messages channel has no buffer and no receiver
		fmt.Println("no message sent")
	}

	// multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
