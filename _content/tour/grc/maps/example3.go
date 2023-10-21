//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ότι μόνο τύποι που μπορούν να έχουν
// ορισμό για την ισότητα μπορούν να είναι κλειδιά σε πίνακες κατακερματισμού.
package main

import "fmt"

// Ο user αναπαριστά κάποιον που χρησιμοποιεί το πρόγραμμα.
type user struct {
	name    string
	surname string
}

// Η users ορίζει ένα σύνολο από user.
type users []user

func main() {

	// Δηλώστε και δημιουργείστε έναν πίνακα κατακερματισμού που χρησιμοποιεί μια φέτα ως κλειδί.
	u := make(map[users]int)

	// ./example3.go:22: invalid map key type users

	// Πραγματοποιείστε επαναληπτική προσπέλαση στην πίνακα κατακερματισμού.
	for key, value := range u {
		fmt.Println(key, value)
	}
}
