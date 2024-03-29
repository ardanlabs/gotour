//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως δηλώνονται και
// χρησιμοποιούνται τύποι συναρτήσεων.
package main

import "fmt"

// Η event παρουσιάζει γεγονότα προσβάσιμα παντού στο πρόγραμμα.
func event(message string) {
	fmt.Println(message)
}

// Ο data είναι ένας struct, στον οποίο είναι δυνατόν να
// δεσμευτούν μέθοδοι τύπου.
type data struct {
	name string
	age  int
}

// Η event παρουσιάζει ένα γεγονός, για αυτά τα data.
func (d *data) event(message string) {
	fmt.Println(d.name, message)
}

// =============================================================================

// Η fireEvent1 χρησιμοποιεί έναν ανώνυμο τύπο συνάρτησης.
func fireEvent1(f func(string)) {
	f("anonymous")
}

// Η handler αναπαριστά μια συνάρτηση για την διαχείριση γεγονότων.
type handler func(string)

// Η fireEvent2 χρησιμοποιεί έναν τύπο συνάρτησης.
func fireEvent2(h handler) {
	h("handler")
}

// =============================================================================

func main() {

	// Δηλώστε μια μεταβλητή τύπου data.
	d := data{
		name: "Bill",
	}

	// Χρησιμοποιήστε τον χειριστή fireEvent1, που αποδέχεται οποιαδήποτε
	// συνάρτηση ή μέθοδο τύπου με την κατάλληλη υπογραφή.
	fireEvent1(event)
	fireEvent1(d.event)

	// Χρησιμοποιήστε τον χειριστή fireEvent2, που αποδέχεται κάθε
	// συνάρτηση ή μέθοδο τύπου `handler` ή κάθε ρητή κατασκευή
	// συνάρτησης ή μεθόδου τύπου, με την κατάλληλη υπογραφή.
	fireEvent2(event)
	fireEvent2(d.event)

	// Δηλώστε μια συνάρτηση τύπου handler για τις συναρτήσεις,
	// που βασίζονται στα γεγονότα, που είναι προσβάσιμα παντού
	// στο πρόγραμμα καθώς και γεγονότα που βασίζονται σε μεθόδους
	// τύπου.
	h1 := handler(event)
	h2 := handler(d.event)

	// Χρησιμοποιήστε τον χειριστή fireEvent2, που αποδέχεται
	// τιμές τύπου handler.
	fireEvent2(h1)
	fireEvent2(h2)

	// Χρησιμοποιήστε τον χειριστή fireEvent1, που αποδέχεται
	// κάθε συνάρτηση ή μέθοδο τύπου, με την σωστή υπογραφή.
	fireEvent1(h1)
	fireEvent1(h2)
}
