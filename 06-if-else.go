package main

import "fmt"

func main() {
	/* NO ternary if (condition ? true-result : false-result) in Go, need full if statement even for basic conditions */

	for i := range 8 {
		if i%2 == 0 { // basic if/else example
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}

		if i%4 == 0 { // can have if without else
			fmt.Println(i, "is divisible by 4")
		}

		if i%2 == 0 || (i+1)%2 == 0 { // can use logical operators like && and ||
			fmt.Println("either", i, "or", i+1, "is even")
		}
	}

	if num := 9; num < 0 { // can declare and initialize variable for scope of current and subsequent branches
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
