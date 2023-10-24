//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί η χρήση του ιδιώματος κόμμα-ok,
// για διαβεβαιώσεις τύπων.
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

// Ο finder αναπαριστά την ικανότητα να αναζητάει κανείς χρήστες.
type finder interface {
	find(id int) (*user, error)
}

// Ο userSVC είναι μια υπηρεσία για την διαχείριση των χρηστών.
type userSVC struct {
	host string
}

// Η find υλοποιεί την διεπαφή finder, χρησιμοποιώντας σημειολογία δείκτη διεύθυνσης.
func (*userSVC) find(id int) (*user, error) {
	return &user{id: id, name: "Anna Walker"}, nil
}

// Ο mockSVC ορίζει μια μίμηση υπηρεσίας, στην οποία θα αποκτήσουμε πρόσβαση.
type mockSVC struct{}

// Η find υλοποιεί την διεπαφή finder, χρησιμοποιώντας σημειολογία δείκτη διεύθυνσης.
func (*mockSVC) find(id int) (*user, error) {
	return &user{id: id, name: "Jacob Walker"}, nil
}

func main() {
	var svc mockSVC

	if err := run(&svc); err != nil {
		log.Fatal(err)
	}
}

func run(f finder) error {
	u, err := f.find(1234)
	if err != nil {
		return err
	}
	fmt.Printf("Found user %+v\n", u)

	// Αν ο τύπος της πραγματικής τιμής, που είναι αποθηκευμένη στην τιμή διεπαφής είναι
	// τύπου *userSVC, τότε η "ok" θα είναι true και η "svc" θα είναι ένα αντίγραφο
	// του δείκτη διεύθυνσης, που είναι αποθηκευμένη στην διεπαφή.
	if svc, ok := f.(*userSVC); ok {
		log.Println("queried", svc.host)
	}

	return nil
}
