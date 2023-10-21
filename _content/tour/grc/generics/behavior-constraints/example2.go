//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος υλοποίησης
// μια λύση κενής διεπαφής η οποία χρησιμοποιεί διαβεβαιώσεις τύπων
// για τις διαφορετικές πραγματικές φέτες που πρέπει να υποστηριχθούν.
// Βασικά, μεταφέραμε τις παραπάνω συναρτήσεις σε δηλώσεις case.
// Αυτή η συνάρτηση χρησιμοποιεί την μέθοδο τύπου String από την τιμή.
package main

import (
	"fmt"
)

func stringifyAssert(v interface{}) []string {
	switch list := v.(type) {
	case []user:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret

	case []customer:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	}

	return nil
}

// Ορίζοντας δύο τύπους που υλοποιούν την διεπαφή the fmt.Stringer. Κάθε
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

	s1 := stringifyAssert(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyAssert(customers)

	fmt.Println("customers:", s2)
}
