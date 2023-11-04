//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως να προσπελαστεί ένας
// σχεσιακός πίνακας κατά αλφαβητική σειρά των κλειδιών του.
package main

import (
	"fmt"
	"sort"
)

// Ο user αναπαριστά κάποιον, που χρησιμοποιεί το πρόγραμμα.
type user struct {
	name    string
	surname string
}

func main() {

	// Δηλώστε και δώστε αρχικές τιμές στον σχεσιακό πίνακα.
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Τραβήξτε τα κλειδιά από τον πίνακα κατακερματισμού.
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	// Ταξινομήστε τα κλειδιά αλφαβητικά.
	sort.Strings(keys)

	// Προσπελάστε τα κλειδιά και τραβήξτε κάθε τιμή,
	// από τον σχεσιακός πίνακα.
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}
