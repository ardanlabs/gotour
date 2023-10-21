//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο συντακτικός κανόνας επιβεβαιώσεων τύπου.
package main

import (
	"fmt"
	"log"
)

// Ο user ορίζει έναν χρήστη στην εφαρμογή μας.
type user struct {
	id   int
	name string
}

// ο finder αναπαριστά την δυνατότητα αναζήτησης χρηστών.
type finder interface {
	find(id int) (*user, error)
}

// Ο userSVC είναι μια υπηρεσία αντιμετώπισης χρηστών.
type userSVC struct {
	host string
}

// Η find υλοποιεί την διεπαφή finder χρησιμοποιώντας σημειολογία δείκτη διεύθυνσης.
func (*userSVC) find(id int) (*user, error) {
	return &user{id: id, name: "Anna Walker"}, nil
}

func main() {
	svc := userSVC{
		host: "localhost:3434",
	}

	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

// Η run πραγματοποιεί την λειτουργία αναζήτησης στα πραγματικά δεδομένα που περνάνε στην κλήση.
func run(f finder) error {
	u, err := f.find(1234)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// Ιδανικά η αφαίρεση finder θα περιελάμβανε όλη την συμπεριφορά
	// που σας ενδιαφέρει. Όμως τι θα συνέβαινε, αν για κάποιο λόγο,
	// χρειάζεται να πάρετε την πραγματική τιμή που είναι αποθηκευμένη
	// στην διεπαφή;

	// Μπορείτε να έχετε πρόσβαση στο πεδίο "host" από τον πραγματικό τύπο
	// δείκτη διεύθυνσης userSVC ο οποίος είναι αποθηκευμένος στην μεταβλητή
	// διεπαφής; Όχι, κάτι τέτοιο δεν είναι δυνατόν, τουλάχιστον όχι άμεσα.
	// Το μόνο που γνωρίζετε είναι ότι τα δεδομένα έχουν μια μέθοδο τύπου με το
	// όνομα "find".
	// ./example5.go:61:26: f.host undefined (type finder has no field or method host)
	log.Println("queried", f.host)

	// Μπορείτε να χρησιμοποιήσετε μια διαβεβαίωση τύπου προκειμένου να πάρετε ένα αντίγραφο
	// του δείκτη διεύθυνσης userSVC που είναι αποθηκευμένος στην διεπαφή.
	svc := f.(*userSVC)
	log.Println("queried", svc.host)

	return nil
}
