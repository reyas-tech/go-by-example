package main

import "fmt"

func main() {
	/* NOTE: slices are more common in Go, arrays are useful in special scenarios */
	var a [5]int           // like "let x: int[5];"
	fmt.Println("emp:", a) // Zero-valued by default

	a[4] = 100
	fmt.Println("set:", a)    // Prints whole array
	fmt.Println("get:", a[4]) // Prints specific element
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} // ... counts number of elements automatically
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // Makes array equal to [100, 0, 0, 400, 500]
	fmt.Println("idx:", b)         // Meaning arrays are 0 based, and unspecified values will be zeroed

	var twoD [2][3]int
	for i := 0; i < 2; i++ { // Array types are one-dimensional
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j // 2D arrays exist, can be defined a number of ways
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d:", twoD)
}
