//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως οι σχεσιακοί πίνακες
// είναι αναφορές.
package main

import "fmt"

func main() {

	// Δώστε αρχικές τιμές στον σχεσιακό πίνακα.
	scores := map[string]int{
		"anna":  21,
		"jacob": 12,
	}

	// Περάστε τον σχεσιακό πίνακα σε μια συνάρτηση, προκειμένου να
	// πραγματοποιήσει κάποια μεταβολή.
	double(scores, "anna")

	// Δείτε, ότι η αλλαγή είναι εμφανής στον σχεσιακό πίνακα μας.
	fmt.Println("Score:", scores["anna"])
}

// Η double βρίσκει το σκορ για συγκεκριμένο παίκτη και το
// πολλαπλασιάζει με το 2.
func double(scores map[string]int, player string) {
	scores[player] = scores[player] * 2
}
