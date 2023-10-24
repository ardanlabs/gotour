//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος υλοποίησης
// μιας λύσης αντανάκλασης η οποία επιτρέπει σε μια φέτα οποιουδήποτε τύπου
// να μετατραπεί σε συμβολοσειρά. Πρόκειται για μια συνάρτηση γενικού
// προγραμματισμού, εξαιτίας του πακέτου reflect. Παρατηρείστε την κλήση
// στην μέθοδο τύπου String μέσω αντανάκλασης.
package main

import (
	"fmt"
	"reflect"
)

func stringifyReflect(v interface{}) []string {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return nil
	}

	ret := make([]string, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		m := val.Index(i).MethodByName("String")
		if !m.IsValid() {
			return nil
		}

		data := m.Call(nil)
		ret = append(ret, data[0].String())
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

	s1 := stringifyReflect(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyReflect(customers)

	fmt.Println("customers:", s2)
}
