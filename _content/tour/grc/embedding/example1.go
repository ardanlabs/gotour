//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως ό,τι κάνουμε ΔΕΝ είναι
// να ενσωματώνουμε ένα τύπο αλλά, απλά, να χρησιμοποιούμε έναν τύπο ως ένα πεδίο.
package main

import "fmt"

// Ο user ορίζει ένα χρήστη στο πρόγραμμα.
type user struct {
	name  string
	email string
}

// Η notify υλοποιεί μια μέθοδο τύπου, που ειδοποιεί τους χρήστες
// για διαφορετικά γεγονότα.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// Ο admin αναπαριστά έναν χρήστη διαχειριστή, με προνόμια.
type admin struct {
	person user // ΔΕΝ είναι Ενσωμάτωση
	level  string
}

func main() {

	// Δημιουργήστε έναν χρήστη admin.
	ad := admin{
		person: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Μπορούμε να έχουμε πρόσβαση σε μεθόδους τύπου των πεδίων.
	ad.person.notify()
}
