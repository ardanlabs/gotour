//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως η χωρητικότητα της φέτας
// δεν είναι διαθέσιμη προς χρήση.
package main

import "fmt"

func main() {

	// Δημιουργείστε μια φέτα με μήκος 5 στοιχεία.
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Δεν μπορείτε να έχετε πρόσβαση σε δείκτη της φέτας, πέραν του μήκους της.
	fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits)
}
