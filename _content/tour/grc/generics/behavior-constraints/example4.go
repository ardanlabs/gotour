//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος υλοποίησης
// μιας λύσης γενικού προγραμματισμού, που επιτρέπει σε μια φέτα κάποιου
// τύπου T (ο οποίος θα καθοριστεί αργότερα) να περάσει και να μετατραπεί
// σε συμβολοσειρά. Αυτός ο κώδικας μοιάζει περισσότερο με τις πραγματικές
// υλοποιήσεις με τις οποίες ξεκινήσαμε και είναι ευκολότερο να διαβαστεί
// από την υλοποίηση της αντανάκλασης. Όμως, εφαρμόζεται ένας περιορισμός
// διεπαφής τύπου fmt.Stringer, προκειμένου να επιτραπεί στον μεταγλωττιστή
// να γνωρίζει ότι η τιμή του τύπου T που περνάει, απαιτεί μια μέθοδο τύπου
// String.
package main

import (
	"fmt"
)

func stringify[T fmt.Stringer](slice []T) []string {
	ret := make([]string, 0, len(slice))

	for _, value := range slice {
		ret = append(ret, value.String())
	}

	return ret
}

// Ορίζοντας δύο τύπους, που υλοποιούν την διεπαφή the fmt.Stringer. Κάθε
// υλοποίησης δημιουργεί μια εκδοχή συμβολοσειράς του πραγματικού τύπου.

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
	name  string
	email string
}

func (u customer) String() string {
	return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

// =============================================================================

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}

	s1 := stringify(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringify(customers)

	fmt.Println("customers:", s2)
}
