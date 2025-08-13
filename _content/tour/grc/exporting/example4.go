//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως μη εξαγόμενα πεδία από 
// έναν εξαγόμενο τύπο struct δεν μπορούν να προσπελαστούν άμεσα.
package main

import (
	"fmt"

	"play.ground/users"
)

func main() {

	// Δημιουργήστε μια τιμή τύπου User από το πακέτο users.
	u := users.User{
		Name: "Chole",
		ID:   10,

		password: "xxxx",
	}

	// ./example4.go:21: unknown field password in struct literal of type users.User

	fmt.Printf("User: %#v\n", u)
}

// -----------------------------------------------------------------------------
-- users/users.go --

// Το πακέτο users παρέχει υποστήριξη για την διαχείριση χρηστών.
package users

// Ο User αναπαριστά πληροφορίες με κάποιο χρήστη.
type User struct {
	Name string
	ID   int

	password string
}

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.25.0
