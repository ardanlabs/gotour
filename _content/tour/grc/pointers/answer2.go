//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε έναν τύπο struct και δημιουργείστε μια τιμή αυτού του τύπου. Δημιουργήστε μια συνάρτηση που
// μπορεί να αλλάξει την τιμή κάποιου πεδίου σε μεταβλητές αυτού του τύπου struct. Παρουσιάστε την τιμή πριν
// και μετά την κλήση της συνάρτησης σας.
package main

import "fmt"

// Ο user αναπαριστά έναν χρήστη στο σύστημα.
type user struct {
	name        string
	email       string
	accessLevel int
}

func main() {

	// Δημιουργείστε μια μεταβλητή τύπου user και δώστε αρχική τιμή σε κάθε πεδίο.
	bill := user{
		name:        "Bill",
		email:       "bill@ardanlabs.com",
		accessLevel: 1,
	}

	// Παρουσιάστε την τιμή της μεταβλητής.
	fmt.Println("access:", bill.accessLevel)

	// Μοιραστείτε την μεταβλητή bill με την συνάρτηση accessLevel
	// μαζί με μια τιμή για να ανανεωθεί με αυτή η τιμή του πεδίου accessLevel.
	accessLevel(&bill, 10)

	// Παρουσιάστε την τιμή του πεδίου accessLevel ξανά.
	fmt.Println("access:", bill.accessLevel)
}

// Η accessLevel αλλάζει την τιμή του επιπέδου πρόσβαση του χρήστη.
func accessLevel(u *user, accessLevel int) {

	// Εκχωρείστε στην τιμή του πεδίου accessLevel field την τιμή
	// που περνάει ως είσοδος της συνάρτησης.
	u.accessLevel = accessLevel
}
