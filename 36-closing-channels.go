package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // special 2-value form of receive, second value is false if jobs has been closed and all values already received
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true // notify done when we've worked on all the jobs
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ { // send 3 jobs to the worker over the jobs channel
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // closing a channel indicates no more values will be sent on it
	fmt.Println("sent all jobs")

	<-done // await the worker using the synchronization approach

	_, ok := <-jobs // closed job returns zero value of underlying type immediately
	fmt.Println("received more jobs:", ok)
}
