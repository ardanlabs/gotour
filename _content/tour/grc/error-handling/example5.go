//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί αν η κλάση μπορεί να εντοπίσει το πρόβλημα.
package main

import (
	"fmt"
	"log"
)

// Ο customError είναι απλά ένας κενός struct.
type customError struct{}

// Η Error υλοποιεί την διεπαφή error.
// Η Error υλοποιεί την διεπαφή error.
func (c *customError) Error() string {
	return "Find the bug."
}

// Η fail επιστρέφει τιμές nil για τους δύο τύπους επιστροφής.
func fail() ([]byte, *customError) {
	return nil, nil
}

func main() {
	var err error
	if _, err = fail(); err != nil {
		log.Fatal("Why did this fail?")
	}

	log.Println("No Error")
}

// =============================================================================

func reason() {
	var err error
	fmt.Printf("Type of value stored inside the interface: %T\n", err)

	if _, err = fail(); err != nil {
		fmt.Printf("Type of value stored inside the interface: %T\n", err)
	}

	log.Println("No Error")
}
