//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Το playground είναι στην πραγματικότητα ένα περιβάλλον 64-bit, με δείκτες διεύθυνσης
// 32-bit.
// Ο συνδυασμός λειτουργικού συστήματος/αρχιτεκτονικής υλικού ονομάζεται nacl/amd64p32

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος δήλωσης μεταβλητών.
package main

import "fmt"

func main() {

	// Δήλωση μεταβλητών που λαμβάνουν την μηδενική τους τιμή.
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Δήλωση μεταβλητών και ανάθεση αρχικών τιμων.
	// Χρήση του τελεστή γρήγορης δήλωσης μεταβλητής.
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Προσδιορισμός τύπου μεταβλητής και πραγματοποίηση μετατροπής.
	aaa := int32(10)

	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}

/*
	Μηδενικές Τιμές:
	Αρχική Τιμή εκ του Τύπου
	Τιμή αληθείας false
	Ακέραιο 0
	Κινητής υποδιαστολής 0
	Μιγαδικό 0i
	Συμβολοσειρά "" (άδεια συμβολοσειρά)
	Δείκτης Διεύθυνσης Μνήμης nil
*/
