//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο εξακρίβωσης αν μια συμβολοσειρά
// είναι μια αντιμετάθεση ή όχι.
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

// Ο RuneSlice είναι ένας προσαρμοσμένος τύπος μιας φέτας ρούνων (rune).
type RuneSlice []rune

// Για την ταξινόμηση ενός RuneSlice.
func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Η IsPermutation ελέγχει αν δύο συμβολοσειρές είναι αντιμεταθέσεις η μια της άλλης.
func IsPermutation(str1, str2 string) bool {

	// Αν δεν έχουν το ίδιο μήκος δεν γίνεται να είναι αντιμεταθέσεις.
	if len(str1) != len(str2) {
		return false
	}

	// Μετατρέψτε κάθε συμβολοσειρά σε μια συλλογή από ρούνους.
	s1 := []rune(str1)
	s2 := []rune(str2)

	// Ταξινομήστε κάθε συλλογή από ρούνους.
	sort.Sort(RuneSlice(s1))
	sort.Sort(RuneSlice(s2))

	// Μετατρέψτε την συλλογή από ρούνους σε μια συμβολοσειρά και
	// συγκρίνετε.
	return string(s1) == string(s2)
}
