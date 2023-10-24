//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Υλοποιήστε έναν τύπο σχεσιακού πίνακα γενικού προγραμματισμού.
package main

import (
	"fmt"
)

// Δηλώστε έναν τύπο γενικού προγραμματισμού με το όνομα
// keymap που χρησιμοποιεί ένα υποκείμενο τύπο σχεσιακό
// πίνακα με κλειδιά τύπου συμβολοσειράς και
// τιμές κάποιου τύπου T.
type keymap[T any] map[string]T

// Υλοποιήστε μια μέθοδο τύπου με το όνομα set, που αποδέχεται
// ένα κλειδί τύπου συμβολοσειράς και μια τιμή τύπου T.
func (km keymap[T]) set(k string, v T) {
	km[k] = v
}

// Υλοποιήστε μια μέθοδο τύπου με το όνομα get που αποδέχεται
// ένα κλειδί τύπου συμβολοσειράς και επιστρέφει μια τιμή τύπου
// T και true ή false αν το κλειδί βρεθεί.
func (km keymap[T]) get(k string) (T, bool) {
	var zero T

	v, found := km[k]
	if !found {
		return zero, false
	}

	return v, true
}

// =============================================================================

func main() {

	// Δημιουργήστε μια τιμή τύπου keymap που αποθηκεύει ακέραιους.
	km := make(keymap[int])

	// Προσθέστε μια τιμή με το κλειδί "jack" και τιμή 10.
	km.set("jack", 10)

	// Προσθέστε μια τιμή με το κλειδί "jill" και τιμή 20.
	km.set("jill", 20)

	// Πάρτε την τιμή για το "jack" και επιβεβαιώστε ότι βρέθηκε.
	jack, found := km.get("jack")
	if !found {
		fmt.Println("jack not found")
		return
	}

	// Τυπώστε την τιμή για το κλειδί "jack".
	fmt.Println("jack", jack)

	// Πάρτε την τιμή για το "jill" και επιβεβαιώστε ότι βρέθηκε.
	jill, found := km.get("jill")
	if !found {
		fmt.Println("jill not found")
		return
	}

	// Τυπώστε την τιμή για το κλειδί "jill".
	fmt.Println("jill", jill)
}
