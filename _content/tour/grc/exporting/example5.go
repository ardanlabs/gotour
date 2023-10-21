//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως δημιουργούνται τιμές από εξαγόμενους τύπους 
// με ενσωματωμένους μη εξαγόμενους τύπους.
package main

import (
	"fmt"

	"play.ground/users"
)

func main() {

	// Δημιουργείστε μια τιμή τύπου Manager Από το πακέτο users.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// Εκχωρείστε τιμή στα εξαγόμενα πεδία από τον μη εξαγόμενο εσωτερικό τύπο χρήστη.
	u.Name = "Chole"
	u.ID = 10

	fmt.Printf("User: %#v\n", u)
}

// -----------------------------------------------------------------------------
-- users/users.go --

// Το πακέτο users παρέχει υποστήριξη για διαχείριση χρηστών.
package users

// Ο User αναπαριστά πληροφορίες σχετικά με κάποιο χρήστη.
type user struct {
	Name string
	ID   int
}

// Ο Manager αναπαριστά πληροφορίες σχετικά με κάποιο manager.
type Manager struct {
	Title string

	user
}

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
