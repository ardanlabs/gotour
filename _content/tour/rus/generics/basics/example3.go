//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to write a function that provides a reflection
// solution which allows a slice of any type to be provided and printed. This
// is a generic function thanks to the reflect package.
package main

import (
	"fmt"
	"reflect"
)

func printReflect(v interface{}) {
	fmt.Print("Reflect: ")

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}

	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3}
	printReflect(numbers)
	print(numbers)

	strings := []string{"A", "B", "C"}
	printReflect(strings)
}
