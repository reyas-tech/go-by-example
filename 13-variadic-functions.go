package main

import "fmt"

/* Variadic functions can be called with any number of arguments (ex: fmt.Println()) */
func sum(nums ...int) { // accepts arbitrary number of ints as arguments (nums is equivalent to []int in the function)
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums { // can iterate over nums with range because it is equivalent to []int in the function
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // if you already have multiple args in a slice, apply them to variadic function using func(slice...)
}
