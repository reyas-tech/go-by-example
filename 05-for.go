package main

import "fmt"

func main() {
	/* for is Go's ONLY looping construct */

	i := 1
	for i <= 3 { // basic type, single condition
		fmt.Println("i", i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // classic initial/condition/after for loop, j only exists in the scope of this loop
		fmt.Println("j", j)
	}

	for i := range 3 { // range over an integer = "do this N times", starts at 0 stops at number after range
		fmt.Println("range", i) // uses i in scope of this for loop, not i from line 8
	}

	j := 0
	for { // without a condition for will loop repeatedly until you break or return
		fmt.Println("loop")
		j++
		if j > 3 {
			break
		}
	}

	for n := range 6 {
		if n%2 == 0 {
			continue // continue to next iteration of loop
		}
		fmt.Println("n", n)
	}
}
