package main

import (
	"errors"
	"fmt"
)

// Possible to use custom types as errors by implementing Error() method on them
type argError struct { // custom error type usually has "Error" suffix
	arg     int
	message string
}

func (e *argError) Error() string { // now argError implements the error interface
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"} // return custom error
	}
	return arg + 3, nil
}

func main() {
	_, err := f2(42)
	var ae *argError
	if errors.As(err, &ae) { // errors.As checks that given error (or any in chain) matches specific error type and converts to a value of that type
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't march argError")
	}
}
