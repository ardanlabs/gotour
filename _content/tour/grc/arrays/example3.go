//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί η συμπεριφορά της έκφρασης for range και
// το πως η μνήμη ενός πίνακα είναι συνεχόμενη.
package main

import "fmt"

func main() {

	// Δηλώστε ένα πίνακα 5 συμβολοσειρών που λαμβάνει κάποιες τιμές ως αρχικές.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Προσπλάστε επανειλημμένα τον πίνακα παρουσιάζοντας την τιμή και την διεύθυνση
	// κάθε στοιχείου.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}
}
