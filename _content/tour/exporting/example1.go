// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// THESE EXAMPLES WON'T RUN SINCE WE NEED PACKAGING

// =============================================================================
// counters/counters.go

// Package counters provides alert counter support.
package counters

// AlertCounter is an exported named type that
// contains an integer counter for alerts.
type AlertCounter int

// =============================================================================

// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"counters"
)

func main() {

	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}
