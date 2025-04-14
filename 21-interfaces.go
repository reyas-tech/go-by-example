package main

import (
	"fmt"
	"math"
)

type geometry interface { // Interfaces are named collections of method signatures
	area() float64
	perim() float64
}

// Implementing interface on rect and circle types
type rect2 struct { // named rect2 since rect is declared in 20-methods.go
	width, height float64
}
type circle struct {
	radius float64
}

// To implement interface, implement all the methods in the interface
// Implement geometry on rects
func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implement geometry on circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	// Call methods that are in the named interface
	fmt.Println("shape", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func detectCircle(g geometry) {
	// Type assertion
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {
	r := rect2{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) // circle and rect2 structs both implement geometry interface
	measure(c) // so we can use instances of these structs as arguments to measure

	detectCircle(r)
	detectCircle(c)
}
