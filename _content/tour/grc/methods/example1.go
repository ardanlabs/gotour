//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως δηλώνονται μέθοδοι τύπου
// και πως τις υποστηρίζει ο μεταγλωττιστής της Go.
package main

import (
	"fmt"
)

// Ο user ορίζει ένα χρήστη στο πρόγραμμα.
type user struct {
	name  string
	email string
}

// Η notify υλοποιεί μια μέθοδο τύπου με δέκτη μεθόδου τιμής.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// Η changeEmail υλοποιεί μια μέθοδο με δέκτη μεθόδου δείκτη διεύθυνσης.
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	// Οι τιμές τύπου user μπορούν να χρησιμοποιηθούν, προκειμένου να
	// κληθούν μέθοδοι τύπου, οι οποίες έχουν δηλωθεί με δέκτες μεθόδου
	// τόσο τιμής όσο και δείκτη διεύθυνσης.
	bill := user{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	bill.notify()

	// Δείκτες διεύθυνσης τύπου user μπορούν να χρησιμοποιηθούν,
	// προκειμένου να κληθούν μέθοδοι τύπου, που είναι δηλωμένες
	// τόσο με δείκτες μεθόδου τιμής όσο και με δείκτη διεύθυνσης.
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	// Δημιουργείστε μια φέτα με τιμές user με δύο user.
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Πραγματοποιείστε επαναληπτική προσπέλαση στην φέτα των user,
	// αλλάζοντας την σημειολογία που χρησιμοποιείτε.
	// Κακή επιλογή!
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Παράδειγμα εξαίρεσης του κανόνα: Χρησιμοποιώντας σημειολογία
	// δείκτη διεύθυνσης για μια συλλογή από συμβολοσειρές.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}
}
