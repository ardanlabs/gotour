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
	counter := counters.AlertCounter(10)

	fmt.Printf("Counter: %d\n", counter)
}

// -----------------------------------------------------------------------------
-- counters/counters.go --

// Package counters provides alert counter support.
package counters

// AlertCounter is an exported named type that
// contains an integer counter for alerts.
type AlertCounter int

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0

replace  "play.ground/counters" => ./counters
