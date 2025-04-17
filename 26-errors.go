package main

import (
	"errors"
	"fmt"
)

/*
 * In Go, we can communicate errors with an explicit separate return value
 * Different than exceptions in languages like Java and Ruby
 * https://pkg.go.dev/errors
 */

func f(arg int) (int, error) { // by convention, errrors are last return value and have built-in type error
	if arg == 42 {
		return -1, errors.New("can't work with 42") // errors.New constructs basic error value with given message
	}
	return arg + 3, nil // nil value for error means no error occurred
}

// Sentinel error: predeclared variable used to sigify specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower) // can wrap errors with higher-level errors to add context
		// wrapped errors create a logical chain that can be queried with functions like errors.Is and errors.As
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil { // common to use in-line errors check in the if line
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) { // errors.Is checks that given error (or any error in its chain) matches specific error value
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Println("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
