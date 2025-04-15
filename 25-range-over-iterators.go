package main

import (
	"fmt"
	"iter"
	"slices"
)

// Redefining from 24-generics.go
type List2[T any] struct {
	head, tail *element2[T]
}
type element2[T any] struct {
	next *element2[T]
	val  T
}

func (lst *List2[T]) Push2(v T) {
	if lst.tail == nil {
		lst.head = &element2[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element2[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// Replace AllElements from 24-generics.go with All
// https://pkg.go.dev/iter#Seq
func (lst *List2[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) { // All returns an ITERATOR; iterator function takes another function as a parameter (called yield by convention)
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return // calls yield for every element we are iterating over, and notes yield's return value for potential early termination
			}
		}
	}
}

// Function returning an iterator over Fibonacci numbers (infinite)
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List2[int]{}
	lst.Push2(10)
	lst.Push2(13)
	lst.Push2(23)

	for e := range lst.All() { // can use range loops on iterators
		fmt.Println(e)
	}

	all := slices.Collect(lst.All()) // Collect takes any iterator and collects all its values into a slice
	fmt.Println("all:", all)

	for n := range genFib() {
		if n >= 10 {
			break // once loop hits break, yield function passed to iterator returns false
		}
		fmt.Println(n)
	}
}
