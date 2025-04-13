package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000 // const statement can appear anywhere a var statement can

	const d = 3e20 / n // perform arithmetic with arbitrary precision
	fmt.Println(d)

	fmt.Println(int64(d)) // numeric constant has no type until it's given one

	fmt.Println(math.Sin(n)) // can be given a type by using it in a context that requires one (variable assignment/function call)
}
