//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic map type.
package main

import (
	"fmt"
)

// Declare a generic type named keymap that uses an underlying type of map
// with a key of type string and a value of some type T.
type keymap[T any] map[string]T

// Implement a method named set that accepts a key of type string and a value
// of type T.
func (km keymap[T]) set(k string, v T) {
	km[k] = v
}

// Implement a method named get that accepts a key of type string and return
// a value of type T and true or false if the key is found.
func (km keymap[T]) get(k string) (T, bool) {
	var zero T

	v, found := km[k]
	if !found {
		return zero, false
	}

	return v, true
}

// =============================================================================

func main() {

	// Construct a value of type keymap that stores integers.
	km := make(keymap[int])

	// Add a value with key "jack" and a value of 10.
	km.set("jack", 10)

	// Add a value with key "jill" and a value of 20.
	km.set("jill", 20)

	// Get the value for "jack" and verify it's found.
	jack, found := km.get("jack")
	if !found {
		fmt.Println("jack not found")
		return
	}

	// Print the value for the key "jack".
	fmt.Println("jack", jack)

	// Get the value for "jill" and verify it's found.
	jill, found := km.get("jill")
	if !found {
		fmt.Println("jill not found")
		return
	}

	// Print the value for the key "jill".
	fmt.Println("jill", jill)
}
