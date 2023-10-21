//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος χρήσης του
// προεγκατεστημένου περιορισμού τύπου "comparable". Μια παράμετρος τύπου
// με τον περιορισμό comparable αποδέχεται ως όρισμα τύπου οποιονδήποτε
// τύπο είναι συγκρίσιμος (comparable). Επιτρέπει την χρήση τψν τελεστών
// == και != με τιμές αυτής της παραμέτρου τύπου.
package main

import "fmt"

func index[T comparable](list []T, find T) int {
	for i, v := range list {
		if v == find {
			return i
		}
	}
	return -1
}

type person struct {
	name  string
	email string
}

func main() {
	durations := []int{5000, 10, 40}
	findDur := 10

	i := index(durations, findDur)
	fmt.Printf("Index: %d for %d\n", i, findDur)

	people := []person{
		{name: "bill", email: "bill@email.com"},
		{name: "jill", email: "jill@email.com"},
		{name: "tony", email: "tony@email.com"},
	}
	findPerson := person{name: "tony", email: "tony@email.com"}

	i = index(people, findPerson)
	fmt.Printf("Index: %d for %s\n", i, findPerson.name)
}
