package main

import "fmt"

/*
 * Go supports anonymous functions, which can form closures
 * Closures allow functions to "capture" and use variables from their surrounding scope, even after that scope has exited
 * A closure is a function value that references variables from outside its body
 * The function can access and modify these variables even after the outer function has finished executing
 */

func intSeq() func() int { // intSeq() returns anonymous function
	i := 0
	return func() int { // returned anonymous function CLOSES OVER variable i to form a closure
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // assign intSeq result (a function) to nextInt, which captures its own i value

	fmt.Println(nextInt()) // i value is updated each time we call nextInt
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq() // state is unique to each particular function
	fmt.Println(newInts())
}
