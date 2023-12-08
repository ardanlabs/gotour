//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to write a function that provides an empty interface
// solution which uses type assertions for the different concrete slices to be
// supported. We've basically moved the functions from above into case statements.
package main

import (
	"fmt"
)

func printAssert(v interface{}) {
	fmt.Print("Assert: ")

	switch list := v.(type) {
	case []int:
		for _, num := range list {
			fmt.Print(num, " ")
		}

	case []string:
		for _, str := range list {
			fmt.Print(str, " ")
		}
	}

	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3}
	printAssert(numbers)

	strings := []string{"A", "B", "C"}
	printAssert(strings)
}
