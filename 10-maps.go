package main

import (
	"fmt"
	"maps"
)

func main() {
	/* Maps are Go's built-in associative data type (aka hashes/dicts) */

	m := make(map[string]int) // use make() to create an empty map
	m["k1"] = 7               // set key/value pairs with name[key] = val syntax
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"] // get value with name[key] syntax
	fmt.Println("v1:", v1)

	v3 := m["k3"] // if key doesn't exist, ZERO VALUE is returned
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // builtin len returns number of key/value pairs

	delete(m, "k2") // builtin delete removes key/value pairs from map based on key
	fmt.Println("map:", m)

	clear(m) // builtin clear removes ALL key/value pairs from map
	fmt.Println("map:", m)

	_, prs := m["k2"]        // optional second return value indicates if key was present in map (used to determine if key is missing or has zero value)
	fmt.Println("prs:", prs) // the _ is a blank identifier, means the first return value is not saved/used

	n := map[string]int{"foo": 1, "bar": 2} // another way to declare and initialize a map
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { // maps package contains utility functions for maps
		fmt.Println("n == n2")
	}
}
