//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare a constraint not based on behavior but
// based on the type of data that is acceptable. This type of constrain is
// important when functions (like Add) need to perform operations (like +)
// that are not supported by all types.
package main

import "fmt"

type addOnly interface {
	string | int | int8 | int16 | int32 | int64 | float64
}

func Add[T addOnly](v1 T, v2 T) T {
	return v1 + v2
}

func main() {
	fmt.Println(Add(10, 20))
	fmt.Println(Add("A", "B"))
	fmt.Println(Add(3.14159, 2.96))
}
