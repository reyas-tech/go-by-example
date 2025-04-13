package main

import (
	"fmt"
	"slices"
)

func main() {
	/* Slices give a more powerful interface to sequences than arrays */
	var s []string // An uninitialized slice equals to nil and has length 0
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)                                  // builtin make creates a slice with non-zero length
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // cap means capacity (different than length, but by default equal to length)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) // Can access/refer to indices same as arrays
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d") // slices support more operations, operations return new slice
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // can copy slices
	fmt.Println("cpy:", c)

	l := s[2:5] // slice[low:high] gives slice of original slice
	fmt.Println("sl1:", l)

	l = s[:5] // slices up to (but excluding) s[5]
	fmt.Println("sl2:", l)

	l = s[2:] // slices up from (and including) s[2]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // can declare and initialize like arrays
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) { // slices package contains utility functions for slices
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // can be composed into multidimensional data structures
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) // length of inner slices can vary (unlike multidimensional arrays)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
