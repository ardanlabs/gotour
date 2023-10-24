//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Υλοποιήστε έναν τύπο στοίβας γενικού προγραμματισμού.
package main

import (
	"errors"
	"fmt"
)

// Δηλώστε ένα τύπο γενικού προγραμματισμού με το όνομα stack, που χρησιμοποιεί
// έναν struct με μοναδικό πεδίο με το όνομα data, που δηλώνεται σαν μια φέτα
// τύπου T.
type stack[T any] struct {
	data []T
}

// Δηλώστε μια μέθοδο τύπου με το όνομα push, που αποδέχεται μια τιμή κάποιου
// τύπου T και προσθέτει την τιμή στην φέτα.
func (s *stack[T]) push(v T) {
	s.data = append(s.data, v)
}

// Δηλώστε μια μέθοδο τύπου με το όνομα pop, που επιστρέφει την τελευταία τιμή
// κάποιου τύπου T, που προστέθηκε στην φέτα και ένα σφάλμα.
func (s *stack[T]) pop() (T, error) {
	var zero T

	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	v := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return v, nil
}

// =============================================================================

func main() {

	// Δημιουργήστε μια τιμή τύπου στοίβας, που αποθηκεύει ακέραιους.
	var s stack[int]

	// Προσθέστε με την Push τις τιμές 10 και 20 στην στοίβα.
	s.push(10)
	s.push(20)

	// Αφαιρέστε με την Pop μια τιμή από την στοίβα.
	v, err := s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Τυπώστε την τιμή που αφαιρέθηκε.
	fmt.Println(v)

	// Αφαιρέστε άλλη τιμή από την στοίβα.
	v, err = s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Τυπώστε την τιμή που αφαιρέθηκε.
	fmt.Println(v)

	// Αφαιρέστε άλλη τιμή από την στοίβα. Κάτι τέτοιο θα πρέπει
	// να επιστρέψει ένα σφάλμα.
	v, err = s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
}
