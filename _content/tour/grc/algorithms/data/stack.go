//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής μιας απλής στοίβας.
package main

import (
	"errors"
	"fmt"
)

func main() {
	const items = 5

	var s Stack

	for i := 0; i < items; i++ {
		name := fmt.Sprintf("Name%d", i)
		s.Push(Data{Name: name})

		fmt.Println("push:", name)
	}

	fmt.Println("------------------")

	f := func(d Data) error {
		fmt.Println("pop:", d.Name)
		return nil
	}

	s.Operate(f)
}

// Ο Data αναπαριστά αυτό που θα αποθηκευτεί στην στοίβα.
type Data struct {
	Name string
}

// Ο Stack αναπαριστά μια στοίβα δεδομένων.
type Stack struct {
	data []Data
}

// Η Make επιτρέπει την δημιουργία μιας στοίβας με
// αρχική χωρητικότητα για αποτελεσματικότητα.
// Διαφορετικά, μια στοίβα μπορεί να χρησιμοποιηθεί
// στην κατάσταση μηδενικής τιμής.
func Make(cap int) *Stack {
	return &Stack{
		data: make([]Data, 0, cap),
	}
}

// Η Count επιστρέφει τον αριθμό των αντικειμένων στην στοίβα.
func (s *Stack) Count() int {
	return len(s.data)
}

// Η προσθέτει δεδομένα στην κορυφή της στοίβας.
func (s *Stack) Push(data Data) {
	s.data = append(s.data, data)
}

// Η Pop απομακρύνει δεδομένα από την κορυφή της στοίβας.
func (s *Stack) Pop() (Data, error) {
	if len(s.data) == 0 {
		return Data{}, errors.New("stack empty")
	}

	// Υπολογείστε τον δείκτη υψηλότερου επιπέδου.
	idx := len(s.data) - 1

	// Αντιγράψτε τα δεδομένα από αυτή την θέση δείκτη.
	data := s.data[idx]

	// Απομακρύνετε τον δείκτη υψηλότερου επιπέδου από την φέτα.
	s.data = s.data[:idx]

	return data, nil
}

// Η Peek παρέχει τα δεδομένα που είναι αποθηκευμένα
// στην στοίβα με βάση το επίπεδο από τον πάτο της στοίβας.
// Μια τιμή 0 θα επιστρέψει το υψηλότερα τοποθετημένο
// κομμάτι δεδομένων.
func (s *Stack) Peek(level int) (Data, error) {
	if level < 0 || level > (len(s.data)-1) {
		return Data{}, errors.New("invalid level position")
	}
	idx := (len(s.data) - 1) - level
	return s.data[idx], nil
}

// Η Operate δέχεται μια συνάρτηση που παίρνει δεδομένα
// και καλεί την συγκεκριμένη συνάρτηση για κάθε κομμάτι
// δεδομένων που βρίσκει. Διατρέχει την στοίβα από πάνω
// προς τα κάτω.
func (s *Stack) Operate(f func(data Data) error) error {
	for i := len(s.data) - 1; i > -1; i-- {
		if err := f(s.data[i]); err != nil {
			return err
		}
	}
	return nil
}
