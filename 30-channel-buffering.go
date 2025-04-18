package main

import "fmt"

/* Channels are UNBUFFERED by default
 * They only accept sends (chan <-) if there is a corresponding receive (<-chan) ready to receive the sent value
 * BUFFERED channels accept limited number of values without a corresponding receiver for those values
 */

func main() {
	messages := make(chan string, 2) // make channel of strings buffering up to 2 values
	// can send values into channel without corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"
	// receive the two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
