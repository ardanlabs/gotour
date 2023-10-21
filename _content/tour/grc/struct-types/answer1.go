//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε έναν τύπο struct προκειμένου να διατηρήσετε πληροφορίες ενός χρήστη (name, email and age).
// Δημιουργήστε μια τιμή αυτού του τύπου, δώστε αρχικές τιμές και παρουσιάστε την τιμή κάθε πεδίου.
//
// Δηλώστε και δώστε αρχική τιμή σε ένα ανώνυμο τύπο struct με τα ίδια τρια πεδία. Παρουσιάστε την τιμή του.
package main

import "fmt"

// Ο user αναπαριστά έναν χρήστη του συστήματος.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Δηλώστε μεταβλητή τύπου user και δώστε αρχική τιμή χρησιμοποιώντας μια ρητή κατασκευή struct.
	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
		age:   45,
	}

	// Παρουσιάστε τις τιμές των πεδίων.
	fmt.Println("Name", bill.name)
	fmt.Println("Email", bill.email)
	fmt.Println("Age", bill.age)

	// Δηλώστε μια μεταβλητή χρησιμοποιώντας μια ανώνυμη struct.
	ed := struct {
		name  string
		email string
		age   int
	}{
		name:  "Ed",
		email: "ed@ardanlabs.com",
		age:   46,
	}

	// Παρουσιάστε τις τιμές των πεδίων.
	fmt.Println("Name", ed.name)
	fmt.Println("Email", ed.email)
	fmt.Println("Age", ed.age)
}
