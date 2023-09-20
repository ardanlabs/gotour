//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to determine if a string is a
// permutation or not.
package main

import (
	"fmt"
	"sort"
)

func main() {
	tt := []struct {
		input1 string
		input2 string
	}{
		{"", ""},
		{"god", "dog"},
		{"god", "do"},
		{"1001", "0110"},
	}

	for _, test := range tt {
		success := IsPermutation(test.input1, test.input2)

		switch success {
		case true:
			fmt.Printf("%q and %q is a permutation\n", test.input1, test.input2)

		case false:
			fmt.Printf("%q and %q is NOT a permutation\n", test.input1, test.input2)
		}
	}
}

// =============================================================================

// RuneSlice a custom type of a slice of runes.
type RuneSlice []rune

// For sorting an RuneSlice.
func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// IsPermutation check if two strings are permutations.
func IsPermutation(str1, str2 string) bool {

	// If the length are not equal they cannot be permutation.
	if len(str1) != len(str2) {
		return false
	}

	// Convert each string into a collection of runes.
	s1 := []rune(str1)
	s2 := []rune(str2)

	// Sort each collection of runes.
	sort.Sort(RuneSlice(s1))
	sort.Sort(RuneSlice(s2))

	// Convert the collection of runes back to a string
	// and compare.
	return string(s1) == string(s2)
}
