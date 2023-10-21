//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος πρόσβασης σε ένα εξαγόμενο αναγνωριστικό.
package main

import (
	"fmt"

	"play.ground/counters"
)

func main() {

	// Δημιουργείστε μια μεταβλητή του εξαγόμενου τύπου και δώστε αρχική τιμή ίση με 10.
	counter := counters.alertCounter(10)

	// ./example2.go:16: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}

// -----------------------------------------------------------------------------
-- counters/counters.go --

// Το πακέτο counters παρέχει υποστήριξη μετρητών προειδοποίησης.
package counters

// Ο alertCounter είναι ένας επώνυμος μη εξαγόμενος τύπος 
// που περιέχει έναν ακέραιο μετρητή για προειδοποιήσεις.
type alertCounter int

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
