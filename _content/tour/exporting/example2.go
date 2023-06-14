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

// =============================================================================

// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"

	"counters"
)

func main() {

	// Create a variable of the unexported type and initialize the value to 10.
	counter := counters.alertCounter(10)

	// ./example2.go:17: cannot refer to unexported name counters.alertCounter
	// ./example2.go:17: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
