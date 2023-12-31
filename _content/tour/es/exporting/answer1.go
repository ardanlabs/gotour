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
	"fmt"

	"play.ground/toy"
)

func main() {

	// Create a value of type toy.
	t := toy.New("Bat", 28)

	// Update the counts.
	t.UpdateOnHand(100)
	t.UpdateSold(2)

	// Display each field separately.
	fmt.Println("Name", t.Name)
	fmt.Println("Weight", t.Weight)
	fmt.Println("OnHand", t.OnHand())
	fmt.Println("Sold", t.Sold())
}

// -----------------------------------------------------------------------------
-- toy/toy.go --

// Package toy contains support for managing toy inventory.
package toy

// Toy represents a toy we sell.
type Toy struct {
	Name   string
	Weight int

	onHand int
	sold   int
}

// New creates values of type toy.
func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

// OnHand returns the current number of this
// toy on hand.
func (t *Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand updates the on hand count and
// returns the current value.
func (t *Toy) UpdateOnHand(count int) int {
	t.onHand += count
	return t.onHand
}

// Sold returns the current number of this
// toy sold.
func (t *Toy) Sold() int {
	return t.sold
}

// UpdateSold updates the sold count and
// returns the current value.
func (t *Toy) UpdateSold(count int) int {
	t.sold += count
	return t.sold
}

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
