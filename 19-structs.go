package main

import "fmt"

/* structs are typed collections of fields; useful for grouping data together to form records */
type person struct {
	name string
	age  int
}

func newPerson(name string) *person { // constructs a new person struct with the given name
	p := person{name: name}
	p.age = 42
	return &p // can safely return a pointer to a local variable
	// will only be cleaned up by garbage collector when there are no active references to it
}

func main() {
	fmt.Println(person{"Bob", 20})              // creates new struct
	fmt.Println(person{name: "Alice", age: 30}) // can name the fields when initializing a struct
	fmt.Println(person{name: "Fred"})           // ommitted values will be zero-valued
	fmt.Println(&person{name: "Ann", age: 40})  // yields a pointer to the struct
	fmt.Println(newPerson("Jon"))               // it's idiomatic to encapsulate new struct creation in constructor functions

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // access struct fields with dot (.)
	sp := &s
	fmt.Println(sp.age) // can use dots with struct pointers - pointers automatically dereferenced
	sp.age = 51         // structs are mutable
	fmt.Println(sp.age)

	// if a struct is only used for a single value, can use an anonymous struct type
	// commonly used for table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
