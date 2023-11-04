//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare a type that is based on the empty
// interface so any value can be added into the vector. Since the last
// function is using the empty interface for the return, users will need to
// perform type assertions to get back to the concrete value being stored
// inside the interface.
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type vectorInterface []interface{}

func (v vectorInterface) last() (interface{}, error) {
	if len(v) == 0 {
		return nil, errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

func main() {
	fmt.Print("vectorInterface : ")

	vItf := vectorInterface{10, "A", 20, "B", 3.14}

	itf, err := vItf.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	switch v := itf.(type) {
	case int:
		if v < 0 {
			fmt.Print("negative integer: ")
		}
	case string:
		if !utf8.ValidString(v) {
			fmt.Print("non-valid string: ")
		}
	default:
		fmt.Printf("unknown type %T: ", v)
	}

	fmt.Printf("value: %v\n", itf)
}
