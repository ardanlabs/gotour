//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare two user defined types based on an
// underlying concrete type. Each type implements a method named last that
// returns the value stored at the highest index position in the vector or an
// error when the vector is empty.
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type vectorInt []int

func (v vectorInt) last() (int, error) {
	if len(v) == 0 {
		return 0, errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

type vectorString []string

func (v vectorString) last() (string, error) {
	if len(v) == 0 {
		return "", errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

func main() {
	fmt.Print("vectorInt : ")

	vInt := vectorInt{10, -1}

	i, err := vInt.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if i < 0 {
		fmt.Print("negative integer: ")
	}

	fmt.Printf("value: %d\n", i)

	// -------------------------------------------------------------------------

	fmt.Print("vectorString : ")

	vStr := vectorString{"A", "B", string([]byte{0xff})}

	s, err := vStr.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}

	fmt.Printf("value: %q\n", s)
}
