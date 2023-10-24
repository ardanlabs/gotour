//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε και δημιουργήστε έναν σχεσιακό πίνακα ακέραιων τιμών, με μια
// συμβολοσειρά ως κλειδί. Προσθέστε στον σχεσιακό πίνακα πέντε τιμές
// και επισκεφθείτε διαδοχικά τα στοιχεία του σχεσιακού πίνακα, για να
// παρουσιάσετε τα ζεύγη κλειδιού/τιμής.
package main

import "fmt"

func main() {

	// Δηλώστε και δημιουργείστε έναν σχεσιακό πίνακα,
	// με τιμές τύπου ακεραίου.
	departments := make(map[string]int)

	// Δώστε μερικές αρχικές τιμές στον σχεσιακό πίνακα.
	departments["IT"] = 20
	departments["Marketing"] = 15
	departments["Executives"] = 5
	departments["Sales"] = 50
	departments["Security"] = 8

	// Παρουσιάστε κάθε ζεύγος κλειδιού/τιμής.
	for key, value := range departments {
		fmt.Printf("Dept: %s People: %d\n", key, value)
	}
}
