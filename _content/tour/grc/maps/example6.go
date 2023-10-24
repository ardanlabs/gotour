//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί η αδυναμία να πάρει κανείς,
// την διεύθυνση μνήμης ενός στοιχείου, του σχεσιακού πίνακα.
package main

// Ο player αναπαριστά κάποιον, που παίζει το παιχνίδι μας.
type player struct {
	name  string
	score int
}

func main() {

	// Δηλώστε έναν σχεσιακό πίνακα, με αρχικές τιμές, χρησιμοποιώντας μια
	// ρητή κατασκευή σχεσιακού πίνακα.
	players := map[string]player{
		"anna":  {"Anna", 42},
		"jacob": {"Jacob", 21},
	}

	// Η προσπάθεια λήψης διεύθυνσης μνήμης ενός στοιχείου του πίνακα,
	// αποτυγχάνει.
	anna := &players["anna"]
	anna.score++

	// ./example4.go:23:10: cannot take the address of players["anna"]

	// Αντ' αυτού, πάρτε το στοιχείο, μετατρέψτε το και βάλτε το πάλι πίσω.
	player := players["anna"]
	player.score++
	players["anna"] = player
}
