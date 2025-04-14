package main

import "fmt"

func plus(a int, b int) int { // function declaration with 2 int arguments that returns an int
	return a + b // Go requires explicit returns
}

func plusPlus(a, b, c int) int { // can omit the type until the final parameter that declares that type
	return a + b + c
}

func main() {
	res := plus(1, 2) // call function like normal
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
