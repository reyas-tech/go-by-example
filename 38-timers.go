package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second) // timers represent a single event in the future

	<-timer1.C // blocks timer's channel C until it sends a value indicating timer fired
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired") // this will never print because timer2 is stopped before it fires
	}()
	stop2 := timer2.Stop() // can cancel a timer before it fires, unlike with time.Sleep()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second) // give timer2 enough time to fire (it won't because it's stopped)
}
