//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως οι πίνακες κατακερματισμού είναι αναφορές.
package main

import "fmt"

func main() {

	// Δώστε αρχικές τιμές στον πίνακα κατακερματισμού.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Περάστε τον πίνακα κατακερματισμού σε μια συνάρτηση προκειμένου να πραγματοποιήσει κάποιοα μεταβολή.
	double(scores, "anna")

	// Δείτε ότι η αλλαγή είναι εμφανής στον πίνακα κατακερματισμού μας.
	fmt.Println("Score:", scores["anna"])
}

// Η double βρίσκειτ το σκορ για συγκεκριμένο παίκτη και το πολλαπλασιάζει με το 2.
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
