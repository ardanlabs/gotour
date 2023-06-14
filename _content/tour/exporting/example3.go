// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// THE EXAMPLE WON'T RUN SINCE WE NEED PACKAGING

// =============================================================================
// counters/counters.go

// Package counters provides alert counter support.
package counters

// alertCounter is an unexported named type that
// contains an integer counter for alerts.
type alertCounter int

// New creates and returns values of the unexported type alertCounter.
func New(value int) alertCounter {
	return alertCounter(value)
}

// =============================================================================

// Sample program to show how the program can access a value
// of an unexported identifier from another package.
package main

import (
	"fmt"

	"counters"
)

func main() {

	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
