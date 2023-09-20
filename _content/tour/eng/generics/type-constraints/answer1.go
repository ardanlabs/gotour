//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic function named copyfy that is constrained to only
// making copies of slices of type string or int.
package main

import (
	"fmt"
)

// Declare an interface named copyer that creates a constraint on
// string and int.
type copyer interface {
	string | int
}

// Implement a generic function named copyfy that accepts a slice of some
// type T but constrained on the copyer interface.
func copyfy[T copyer](src []T) []T {
	dest := make([]T, len(src))

	copy(dest, src)

	return dest
}

// =============================================================================

func main() {

	// Construct a slice of string with three values.
	src1 := []string{"Bill", "Jill", "Joan"}

	// Call the copyfy function to make a copy of the slice.
	dest1 := copyfy(src1)

	// Display the slice and the copy.
	fmt.Println("src string :", src1)
	fmt.Println("dest string:", dest1)

	// -------------------------------------------------------------------------

	// Construct a slice of int with three values.
	src2 := []int{10, 20, 30}

	// Call the copyfy function to make a copy of the slice.
	dest2 := copyfy(src2)

	// Display the slice and the copy.
	fmt.Println("src int :", src2)
	fmt.Println("dest int:", dest2)
}
