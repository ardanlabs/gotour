//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε ένα struct που αντιπροσωπεύει έναν παίκτη baseball. Συμπεριλάβετε τα πεδία name,
// atBats and hits.
// Δηλώστε μια μέθοδο τύπου που υπολογίζει την μέση τιμή των χτυπημάτων ενός παίκτη. Ο τύπος
// είναι Hits / AtBats.
// Δηλώστε μια φέτα αυτού του τύπου και δώστε ως αρχική τιμή αρκετούς παίκτες. Πραγματοποιείστε
// διαδοχική επίσκεψη των στοιχείων της φέτας, παρουσιάζοντας στην συνέχεια το πεδίο name και την
// μέση τιμή χτυπημάτων του κάθε παίκτη.
package main

import "fmt"

// Ο player αναπαριστά ένα πρώσωπο στο παιχνίδι.
type player struct {
	name   string
	atBats int
	hits   int
}

// Η average υπολογίζει τον μέσο όρο των χτυπημάτων για ένα παίκτη.
func (p *player) average() float64 {
	if p.atBats == 0 {
		return 0.0
	}

	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Δημιουργήστε μερικούς παίκτες.
	ps := []player{
		{"bill", 10, 7},
		{"jim", 12, 6},
		{"ed", 6, 4},
	}

	// Παρουσιάστε τον μέσο όρο χτυπημάτων για κάθε παίκτη.
	for i := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", ps[i].name, ps[i].average()*1000)
	}

	// Γιατί δεν επέλεξα αυτή την μορφή;
	for _, p := range ps {
		fmt.Printf("%s: AVG[.%.f]\n", p.name, p.average()*1000)
	}
}
