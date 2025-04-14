package main

import "fmt"

/* Pointers allow you to pass references to values/records within your program */

func zeroval(ival int) {
	ival = 0 // only changing ival in scope of this function, doesn't change value of the argument passed
}

func zeroptr(iptr *int) {
	*iptr = 0 // assigning value to dereferenced pointer changes the value at the referenced address
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i) // doesn't change i value
	fmt.Println("zeroval:", i)

	zeroptr(&i) // does change i value
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // &i syntax gives memory address of i (aka pointer to i)
}
