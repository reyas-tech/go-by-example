package main

import "fmt"

/*
 * Go v1.18+ has support for GENERICS, aka type parameters
 * See https://go.dev/blog/deconstructing-type-parameters for more info
 */

/*
 * This function exists in slices library as slices.Index
 * Takes a slice of any comparable type and an element of that type and returns the first occurence of v in s
 * comparable constraint means you compare values of that type with == and != operators
 */
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct { // generic type example
	head, tail *element[T] // singly-linked list with values of any type
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) { // can define methods on generic types just like regular types
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T { // returns all list elements as a slice
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo (implicit):", SlicesIndex(s, "zoo"))                   // compiler infers types automatically (type inference)
	fmt.Println("index of zoo (explicit):", SlicesIndex[[]string, string](s, "zoo")) // can specify types explicitly, but don't need to

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
