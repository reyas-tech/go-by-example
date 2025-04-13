package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i { // basic switch
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // Use commas to separate multiple expressions in same case
		fmt.Println("It's the weekend")
	default: // optional, like an "else" statement
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // switch with no expression = alternative to if/else
	case t.Hour() < 12: // can have conditions for cases instead of constants
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) { // a type switch compares types instead of values
		case bool:
			fmt.Println(t, "is a boolean")
		case int:
			fmt.Println(t, "is an integer")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
