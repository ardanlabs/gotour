//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, παρουσιάζει πως οι διαβεβαιώσεις τύπου είναι
// κατασκευές σταδίου εκτέλεσης και όχι σταδίου μεταγλώττισης.
package main

import (
	"fmt"
	"math/rand"
)

// Ο car αναπαριστά κάτι που οδηγείται.
type car struct{}

// Η String υλοποιεί την διεπαφή fmt.Stringer.
func (car) String() string {
	return "Vroom!"
}

// Ο cloud αναπαριστά ένα μέρος, που αποθηκεύει κανείς πληροφορίες.
type cloud struct{}

// Η String υλοποιεί την διεπαφή fmt.Stringer.
func (cloud) String() string {
	return "Big Data!"
}

func main() {

	// Δημιουργείστε μια φέτα τιμών της διεπαφής Stringer.
	mvs := []fmt.Stringer{
		car{},
		cloud{},
	}

	// Ας τρέξουμε αυτό το παράδειγμα δέκα φορές.
	for i := 0; i < 10; i++ {

		// Επιλέξτε έναν τυχαίο αριθμό μεταξύ 0 και 1.
		rn := rand.Intn(2)

		// Πραγματοποιείστε μια διαβεβαίωση τύπου, προκειμένου να
		// εξακριβωθεί ότι έχουμε έναν πραγματικό τύπο
		// cloud στην τιμή διεπαφής, που επιλέξαμε τυχαία.
		if v, is := mvs[rn].(cloud); is {
			fmt.Println("Got Lucky:", v)
			continue
		}

		fmt.Println("Got Unlucky")
	}
}
