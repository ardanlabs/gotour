//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to write a function that provides a generics
// solution which allows a slice of any type T (to be determined later) to be
// passed and printed.
package main

import (
	"fmt"
)

// To avoid the ambiguity with array declarations, type parameters require a
// constraint to be applied. The `any` constraint states there is no constraint
// on what type T can become. The predeclared identifier `any` is an alias for
// `interface{}`.
//
// This code more closely resembles the concrete implementations that we started
// with and is easier to read than the reflect implementation.

func print[T any](slice []T) {
	fmt.Print("Generic: ")

	for _, v := range slice {
		fmt.Print(v, " ")
	}

	fmt.Print("\n")
}

// =============================================================================

func main() {
	numbers := []int{1, 2, 3}
	print(numbers)

	strings := []string{"A", "B", "C"}
	print(strings)
}
