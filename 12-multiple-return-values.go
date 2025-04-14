package main

import "fmt"

/* Multiple return types often used in idiomatic Go (ex: return result and error values from a function) */

func vals() (int, int) { // 2 return types
	return 3, 7
}

func main() {
	a, b := vals() // multiple assignments for multiple returns
	fmt.Println(a, b)

	_, c := vals() // can use blank identifier if you only want a subset of the returns
	fmt.Println(c)
}
