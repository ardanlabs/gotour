//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
package main

import (
	"play.ground/toy"
)

func main() {

	// Use the New function from the toy package to create a value of
	// type toy.

	// Use the methods from the toy value to set some initialize
	// values.

	// Display each field separately from the toy value.
}

// -----------------------------------------------------------------------------
-- toy/toy.go --

// Package toy contains support for managing toy inventory.
package toy

// Declare a struct type named Toy with four fields. Name string,
// Weight int, onHand int and sold int.

// Declare a function named New that accepts values for the
// exported fields. Return a pointer of type Toy that is initialized
// with the parameters.

// Declare a method named OnHand with a pointer receiver that
// returns the current on hand count.

// Declare a method named UpdateOnHand with a pointer receiver that
// updates and returns the current on hand count.

// Declare a method named Sold with a pointer receiver that
// returns the current sold count.

// Declare a method named UpdateSold with a pointer receiver that
// updates and returns the current sold count.

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
