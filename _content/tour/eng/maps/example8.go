//go:build OMIT || norun

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how nil maps works with set operation.
package main

import "fmt"

// user represents someone using the program.
type user struct {
	name    string
	surname string
}

func main() {
	var users map[string]user

	value, exists := users["Roy"] // safe operation
	fmt.Println(value)            // empty struct
	fmt.Println(exists)           // false

	users["Roy"] = user{"Rob", "Roy"} // panic: assignment to entry in nil map
}
