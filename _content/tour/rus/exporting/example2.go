//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to access an exported identifier.
package main

import (
	"fmt"

	"play.ground/counters"
)

func main() {

	// Create a variable of the exported type and initialize the value to 10.
	counter := counters.alertCounter(10)

	// ./example2.go:16: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}

// -----------------------------------------------------------------------------
-- counters/counters.go --

// Package counters provides alert counter support.
package counters

// alertCounter is an unexported named type that
// contains an integer counter for alerts.
type alertCounter int

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
