//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare a generics version of the user defined
// type. It represents a slice of some type T (to be determined later). The
// last method is declared with a value receiver based on a vector of the
// same type T and returns a value of that same type T as well.
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type vector[T any] []T

func (v vector[T]) last() (T, error) {
	var zero T

	if len(v) == 0 {
		return zero, errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

func main() {
	fmt.Print("vector[int] : ")

	vGenInt := vector[int]{10, -1}

	i, err := vGenInt.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if i < 0 {
		fmt.Print("negative integer: ")
	}

	fmt.Printf("value: %d\n", i)

	// -------------------------------------------------------------------------

	fmt.Print("vector[string] : ")

	vGenStr := vector[string]{"A", "B", string([]byte{0xff})}

	s, err := vGenStr.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}

	fmt.Printf("value: %q\n", s)
}
