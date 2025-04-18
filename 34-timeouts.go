package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1) // channel is buffered, aka nonblocking; common pattern to prevent goroutine leaks if channel never read
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select { // select proceeds with first receive that is ready
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1") // this prints because timeout is received before c1
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res) // this prints because c2 is received before timeout
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
