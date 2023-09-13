//go:build OMIT

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

	// Convert the input string into slice of runes for processing.
	// A rune represent a code point in the UTF-8 character set.
	runes := []rune(input)

	// Run over runes forward and backward comparing runes.
	// If runes[i] != runes[len(runes)-i-1] then it's not a palindrome.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
