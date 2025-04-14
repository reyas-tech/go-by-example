package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // calls itself until it reaches base case of fact(0)
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int // anonymous functions can be recursive
	// but requires explicitly declaring a variable with var to store the function before it's defined
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2) // since fib was declared in main, Go knows which function to call with fib
	}
	fmt.Println(fib(7))
}
