//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος δήλωσης
// δύο τύπων ορισμένων από τον χρήστη βασισμένων σε έναν υποκείμενο
// πραγματικό (concrete) τύπο. Κάθε τύπος υλοποιεί μια μέθοδο τύπου με το
// όνομα last που επιστρέφει την τιμή που είναι αποθηκευμένη στην θέση
// υψηλότερου δείκτη στο διάνυσμα ή ένα σφάλμα όταν το διάνυσμα είναι άδειο.
package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type vectorInt []int

func (v vectorInt) last() (int, error) {
	if len(v) == 0 {
		return 0, errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

type vectorString []string

func (v vectorString) last() (string, error) {
	if len(v) == 0 {
		return "", errors.New("empty")
	}

	return v[len(v)-1], nil
}

// =============================================================================

func main() {
	fmt.Print("vectorInt : ")

	vInt := vectorInt{10, -1}

	i, err := vInt.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if i < 0 {
		fmt.Print("negative integer: ")
	}

	fmt.Printf("value: %d\n", i)

	// -------------------------------------------------------------------------

	fmt.Print("vectorString : ")

	vStr := vectorString{"A", "B", string([]byte{0xff})}

	s, err := vStr.last()
	if err != nil {
		fmt.Print("ERROR:", err)
		return
	}

	if !utf8.ValidString(s) {
		fmt.Print("non-valid string: ")
	}

	fmt.Printf("value: %q\n", s)
}
