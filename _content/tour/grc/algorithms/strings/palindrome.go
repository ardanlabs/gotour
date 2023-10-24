//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό είναι ένα δείγμα προγράμματος, που παρουσιάζει τον τρόπο
// προσδιορισμού μιας παλίνδρομης, από μια κανονική συμβολοσειράς.
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

// Η IsPalindrome ελέγχει αν μια συμβολοσειρά είναι Παλίνδρομη.
func IsPalindrome(input string) bool {

	// Αν η συμβολοσειρά εισόδου είναι άδεια ή έχει μήκος 1,
	// επιστρέψτε true.
	if input == "" || len(input) == 1 {
		return true
	}

	// Μετατρέψτε την συμβολοσειρά εισόδου σε φέτα ρούνων (στμ. rune)
	// για επεξεργασία. Ένας ρούνος αναπαριστά ένα στοιχείο κωδικοποίησης
	// στο σύνολο χαρακτήρων UTF-8.
	runes := []rune(input)

	// Προσπελάστε τους ρούνους κανονικά και ανάποδα, συγκρίνοντας τους.
	// Αν runes[i] != runes[len(runes)-i-1] τότε δεν είναι παλίνδρομη.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
