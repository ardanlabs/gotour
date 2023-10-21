//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί η βασική έννοια της χρήσης δείκτη διεύθυνσης μνήμης
// προκειμένου να μοιραστούν δεδομένα.
package main

import "fmt"

// Ο user αναπαριστά έναν χρήστη στο σύστημα.
type user struct {
	name   string
	email  string
	logins int
}

func main() {

	// Δηλώστε και εκχωρείστε αρχική τιμή σε μια μεταβλητή με το όνομα bill τύπου user.
	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	//** Δεν είναι ανάγκη να συμπεριλάβουμε όλα τα πεδία όταν ορίζουμε τιμές για τα ονόματα των πεδίων
	// με μια ρητή κατασκευή struct.

	// Περάστε την "διεύθυνση της" τιμής bill.
	display(&bill)

	// Περάστε την "διεύθυνση του" πεδίου logins από την τιμή bill.
	increment(&bill.logins)

	// Περάστε την "διεύθυνση της" τιμής bill.
	display(&bill)
}

// Η increment δηλώνει την παράμετρο logins ως μια μεταβλητή δείκτη διεύθυνσης η τιμή της οποίας είναι
// πάντα μια διεύθυνση και δείχνει σε τιμές τύπου int.
func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

// Η display δηλώνει την παράμετρο u ως μια μεταβλητή δείκτη διεύθυνσης τύπου user της οποίας η τιμή είναι πάντα μια
// διεύθυνση μνήμης και δείχνει σε τιμές τύπου user.
func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
}
