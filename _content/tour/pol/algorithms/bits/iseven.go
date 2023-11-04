//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to check if an integer is even or
// odd using bit manipulation.
package main

import (
	"fmt"
)

func main() {

	fmt.Println(8, ":", IsEven(8))
	fmt.Println(15, ":", IsEven(15))
	fmt.Println(4, ":", IsEven(4))
}

// IsEven checks is an integer is even.
func IsEven(num int) bool {

	// Use the bitwise AND operator to see if the least significant
	// bit (LSB) is 0.

	// Helpful source: https://catonmat.net/low-level-bit-hacks
	// 0 & 1 = 0 (even number)
	// 1 & 1 = 1 (odd number)

	return num&1 == 0
}
