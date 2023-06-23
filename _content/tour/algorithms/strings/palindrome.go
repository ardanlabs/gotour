//go:build OMIT
// +build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to determine if a string is a
// palindrome or not.
package main

import "fmt"

func main() {
	tt := []string{"", "G", "bob", "otto", "汉字汉", "test"}

	for _, input := range tt {
		success := IsPalindrome(input)

		switch success {
		case true:
			fmt.Printf("%q is a palindrome\n", input)

		case false:
			fmt.Printf("%q is NOT a palindrome\n", input)
		}
	}
}

// =============================================================================

// IsPalindrome checks if a string is a Palindrome.
func IsPalindrome(input string) bool {

	// If the input string is empty or as a length of 1 return true.
	if input == "" || len(input) == 1 {
		return true
	}

	// Create a reverse string from input string.
	rev := reverseString(input)

	// Check if input and rev strings are equal.
	if input == rev {
		return true
	}

	return false
}

// reverseString takes the specified string and reverses it.
func reverseString(str string) string {

	// Convert the input string into slice of runes for processing.
	// A rune represent a code point in the UTF-8 character set.
	runes := []rune(str)

	// Create an index that will traverse the collection of
	// runes from the beginning to the end.
	var beg int

	// Create an index that will traverse the collection of
	// runes from the end to the beginning.
	end := len(runes) - 1

	// Keep swapping runes until the two indexes meet in the middle.
	for beg < end {

		// Swap the position of these two rune’s.
		r := runes[beg]
		runes[beg] = runes[end]
		runes[end] = r

		// Move the indexes closer to each other
		// working towards the middle of the collection.
		beg = beg + 1
		end = end - 1
	}

	return string(runes)
}
