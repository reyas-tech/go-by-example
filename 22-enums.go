package main

import "fmt"

/*
 * An enum is a type that has a fixed number of possible values, each with a distinctive name
 * Go doesn't have an enum type but enums are simple to implement using existing language idioms
 */

type ServerState int // enum type ServerState has an underlying int type

const ( // possible values for ServerState are defined as constants
	StateIdle ServerState = iota // keyword iota generates successive constant values automatically
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string { // implementing fmt.Stringer interface to print/convert ServerState values
	return stateName[ss] // if many possible values, can use stringer tool in conjuction with go:generate to automate process
	// see https://pkg.go.dev/golang.org/x/tools/cmd/stringer for example
}

func main() {
	ns := transition(StateIdle) // can't pass int type to transition, compiler will complain about type mismatch
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

func transition(s ServerState) ServerState { // emulates state transition for a server
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
