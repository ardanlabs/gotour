//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος υλοποίησης μιας
// συνάρτηση stringify, που είναι συγκεκριμένη σε κάθε ένα από τους πραγματικούς
// τύπους, που είναι υλοποιημένοι παραπάνω. Σε κάθε περίπτωση η συνάρτηση
// stringify επιστρέφει μια φέτα από συμβολοσειρές. Αυτές οι συναρτήσεις
// χρησιμοποιούν την μέθοδο τύπου String σε κάθε ξεχωριστή τιμή user ή
// customer.
package main

import (
	"fmt"
)

func stringifyUsers(users []user) []string {
	ret := make([]string, 0, len(users))
	for _, user := range users {
		ret = append(ret, user.String())
	}
	return ret
}

func stringifyCustomers(customers []customer) []string {
	ret := make([]string, 0, len(customers))
	for _, customer := range customers {
		ret = append(ret, customer.String())
	}
	return ret
}

// Ορίζοντας δύο τύπους, που υλοποιούν την διεπαφή fmt.Stringer. Κάθε
// υλοποίηση δημιουργεί μια εκδοχή συμβολοσειράς του πραγματικού τύπου.

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

	s1 := stringifyUsers(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyCustomers(customers)

	fmt.Println("customers:", s2)
}
