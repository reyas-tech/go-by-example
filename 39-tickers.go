package main

import (
	"fmt"
	"time"
)

/*
 * Timers are for when you want to do something once in the future
 * TICKERS are for when you want to do something repeatedly at regular intervals
 */

func main() {
	ticker := time.NewTicker(500 * time.Millisecond) // tickers are similar to timers: a channel that is sent values
	done := make(chan bool)

	go func() {
		for {
			select { // use built-in select to await the values as they arrive every 500ms
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // tickers can be stopped like timers, won't receive any more values on its channel
	done <- true
	fmt.Println("Ticker stopped")
}
