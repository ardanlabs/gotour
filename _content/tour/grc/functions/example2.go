//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκιεμένου να παρουσιαστεί ο τρόπος με τον οποίο μπορεί κανείς
// να χρησιμοποιείσει το κενό αναγνωριστικό προκειμένου να αγνοήσει τιμές επιστροφής.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Ο user είναι ένας τύπος struct που δηλώνει πληροφορίες χρήστη.
type user struct {
	ID   int
	Name string
}

// Η updateStats παρέχει στατιστικά στοιχεία ανανέωσης.
type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {

	// Δηλώστε και εκχωρείστε αρχική τιμή σε τιμή τύπου user.
	u := user{
		ID:   1432,
		Name: "Betty",
	}

	// Ανανεώστε το πεδίο Name του user. Δεν σας απασχολούν τα στατιστικά στοιχεία ανανέωσης.
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}

	// Παρουσιάστε ότι η ανανέωση ήταν επιτυχημένη.
	fmt.Println("Updated user record for ID", u.ID)
}

// Η updateUser ανανενώνει το συγκεκριμένο έγραφο του user.
func updateUser(u *user) (*updateStats, error) {

	// Η response μιμείται μια απάντηση σε μορφή JSON.
	response := `{"Modified":1, "Duration":0.005, "Success" : true, "Message": "updated"}`

	// Αποσειριοποιείστε (unmarshal) το έγγραφο json σε μια τιμή
	// τύπου userStats struct.
	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	// Ελέγξτε την κατάσταση ανανέωσης προκειμένου να επιβεβαιώσετε ότι η ανανέωση
	// ήταν πετυχημένη.
	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
