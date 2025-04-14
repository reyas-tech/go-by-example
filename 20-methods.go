package main

import "fmt"

/*
 * Go supports METHODS defined on struct types
 * Methods can be defined for either pointer or value receiver types
 */
type rect struct {
	width, height int
}

func (r *rect) area() int { // can use pointer receiver type to avoid copying on method calls or allow method to mutate receiving struct
	return r.width * r.height
}

func (r rect) perim() int { // the perim method has a value receiver type of rect
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	// call the two methods defined for rect struct
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())
	// Go automatically handles conversion between values and pointers for method calls
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
