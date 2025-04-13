package main

import "fmt"

func main() {
	var a = "initial" // var declares one or more variables
	fmt.Println("a:", a)

	var b, c int = 1, 2 // can declare and initialize multiple variables at once
	fmt.Println("b, c:", b, c)

	var d = true
	fmt.Println("d:", d)

	var e int // variables with no initialization are ZERO-VALUED
	fmt.Println("e:", e)
	var e2 string // not a number, EMPTY string
	fmt.Println("e2:", e2)

	f := "apple" // shorthand for var f string = "apple"
	fmt.Println("f:", f)
}
