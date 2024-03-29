//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε έναν τύπο struct, που διατηρεί πληροφορίες για ένα χρήστη. Δηλώστε μια
// συνάρτηση, που δημιουργεί μια τιμή και επιστρέφει δείκτες διεύθυνσης αυτού του
// τύπου και μια τιμή σφάλματος. Καλέστε αυτή την συνάρτηση από την main και
// παρουσιάστε την τιμή.
//
// Κάντε μια δεύτερη κλήση στην συνάρτηση σας, όμως αυτή τη φορά αγνοείστε την τιμή
// και απλά ελέγξτε την τιμή σφάλματος.
package main

import "fmt"

// Ο user αναπαριστά έναν χρήστη του συστήματος.
type user struct {
	name  string
	email string
}

// Η συνάρτηση newUser δημιουργεί και επιστρέφει δείκτες διεύθυνσης τιμών τύπου user.
func newUser() (*user, error) {
	return &user{"Bill", "bill@ardanlabs.com"}, nil
}

func main() {

	// Δημιουργήστε μια τιμή τύπου user.
	u, err := newUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Παρουσιάστε την τιμή.
	fmt.Println(*u)

	// Καλέστε την συνάρτηση και απλά ελέγξτε το σφάλμα στην επιστροφή.
	_, err = newUser()
	if err != nil {
		fmt.Println(err)
		return
	}
}
