//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως να ενσωματωθεί ένας τύπος σε ένα άλλο τύπο καθώς
// και να παρουσιαστεί η σχέση μεταξύ του εσώτερου και του εξώτερου τύπου.
package main

import "fmt"

// Ο user ορίζει έναν χρήστη στο πρόγραμμα.
type user struct {
	name  string
	email string
}

// Η notify υλοποιεί μια μέθοδο τύπου που ενημερώνει τους χρήστες
// για διαφορετικά γεγονότα.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// Ο admin αναπαριστά έναν χρήστη διαχειριστή με προνόμια.
type admin struct {
	user  // Embedded Type
	level string
}

func main() {

	// Δημιουργείστε έναν χρήστη διαχειριστή.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Μπορούμε να έχουμε πρόσβαση στην μέθοδο τύπου του εσώτερου τύπου απευθείας.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}
