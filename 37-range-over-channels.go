package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) // can close non-empty channel and still receive remaining values
	// if channel is not closed, execution results in deadlock error
	for elem := range queue { // iterate over each element in queue, then terminate
		fmt.Println(elem)
	}
}
