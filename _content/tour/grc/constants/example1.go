//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, που παρουσιάζει πως δηλώνονται σταθερές και
// η υλοποίηση τους στην Go.
package main

import "fmt"

func main() {

	// Οι σταθερές ζουν εντός του μεταγλωττιστή.
	// Έχουν ένα παράλληλο σύστημα τύπων.
	// Ο μεταγλωττιστής μπορεί να πραγματοποιήσει σιωπηρές μετατροπές των σταθερών
	// χωρίς τύπο.

	// Σταθερές Χωρίς Τύπο.
	const ui = 12345    // είδος: ακέραιος
	const uf = 3.141592 // είδος: κινητής υποδιαστολής

	// Οι σταθερές με τυπο χρησιμοποιούν το σύστημα τύπων των σταθερών, όμως η
	// ακρίβειά τους είναι περιορισμένη.
	const ti int = 12345        // τύπος: int
	const tf float64 = 3.141592 // τύπος: float64

	// ./constants.go:XX: η σταθερά 1000 υπερχειλίζει έναν uint8
	// const myUint8 uint8 = 1000

	// Η αριθμητική των σταθερών υποστηρίζει διαφορετικά είδη.
	// Η προώθηση είδους χρησιμοποιείται, προκειμένου να προσδιοριστεί το είδος
	// σε τέτοια σενάρια.

	// Η μεταβλητή answer θα είναι τύπου float64.
	var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)
	fmt.Println(answer)

	// Η σταθερά τρίτο θα είναι είδους κινητής υποδιαστολής
	const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)
	fmt.Println(third)

	// Η σταθερά μηδέν θα είναι είδους ακεραίου.
	const zero = 1 / 3 // KindInt(1) / KindInt(3)
	fmt.Println(zero)

	// Αυτό είναι ένα παράδειγμα αριθμητικής με σταθερές, μεταξύ σταθερών με
	// και χωρίς τύπους. Πρέπει να υπάρχουν ίδιοι τύποι, προκειμένου να γίνουν
	// πράξεις.
	const one int8 = 1
	const two = 2 * one // int8(2) * int8(1)
	fmt.Println(two)
}
