package main

import "fmt"

/* range interates over elements in variety of built-in data structures */

func main() {
	// Use range to sum numbers in a slice (arrays work this way too)
	nums := []int{2, 3, 4}
	sum := 0
	// range on arrays/slices provides index and value for each entry
	for _, num := range nums { // blank identifier for index value because we don't need it
		fmt.Println(num)
		sum += num
	}
	fmt.Println("sum:", sum)
	for i, num := range nums { // using both index and value
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// Use range on maps
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// range on map iterates over key/value pairs
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// can also iterate over just keys of map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	for i, c := range "go" { // first value is starting byte index of the rune, second is rune itself
		fmt.Println(i, c)
	}
}
