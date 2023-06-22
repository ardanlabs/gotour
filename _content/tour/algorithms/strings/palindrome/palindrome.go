//go:build OMIT
// +build OMIT

package main

func main() {

}

// Is checks if a string is a Palindrome.
func Is(input string) bool {

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
