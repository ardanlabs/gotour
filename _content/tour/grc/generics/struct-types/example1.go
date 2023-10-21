//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος δήλωσης τύπου γενικού προγραμματισμού με την χρήση τύπου struct.
package main

import (
	"fmt"
)

// Ο κώδικας ορίζει δύο τύπους ορισμένους από τον χρήστη που υλοποιούν μια
// συνδεδεμένη λίστα. Ο τύπος στοιχείου της λίστας περιέχει δεδομένα κάποιου
// τύπου T (που θα καθοριστεί αργότερα) και δείχνει προς άλλα στοιχεία του
// ίδιου τύπου T. Ο τύπος της λίστας περιέχει δείκτες διεύθυνσης στο πρώτο
// και στο τελευταίο στοιχείο κάποιου τύπου T. Η μέθοδος τύπου δηλώνεται
// με λήπτη μεθόδου δείκτη διεύθυνσης με βάση μια λίστα κάποιου τύπου T και
// υλοποιείται προκειμένου να προσθέτει στοιχεία στην λίστα ίδιου τύπου T.

type node[T any] struct {
	Data T
	next *node[T]
	prev *node[T]
}

type list[T any] struct {
	first *node[T]
	last  *node[T]
}

func (l *list[T]) add(data T) *node[T] {
	n := node[T]{
		Data: data,
		prev: l.last,
	}
	if l.first == nil {
		l.first = &n
		l.last = &n
		return &n
	}
	l.last.next = &n
	l.last = &n
	return &n
}

// Αυτός ο τύπος user αναπαριστά τα δεδομένα που θα αποθηκευτούν στην συνδεδεμένη λίστα.

type user struct {
	name string
}

// =============================================================================

func main() {

	// Αποθηκεύστε τιμές τύπου user στην λίστα.
	var lv list[user]
	n1 := lv.add(user{"bill"})
	n2 := lv.add(user{"ale"})
	fmt.Println(n1.Data, n2.Data)

	// Αποθηκεύστε δείκτες διεύθυνσης τύπου user στην λίστα.
	var lp list[*user]
	n3 := lp.add(&user{"bill"})
	n4 := lp.add(&user{"ale"})
	fmt.Println(n3.Data, n4.Data)
}
