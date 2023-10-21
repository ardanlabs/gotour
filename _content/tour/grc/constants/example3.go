//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Παράδειγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος λειτουργίας του iota.
package main

import "fmt"

func main() {

	const (
		A1 = iota // 0 : Ξεκινάει στο 0
		B1 = iota // 1 : Αυξάνει κατά 1
		C1 = iota // 2 : Αυξάνει κατά 1
	)

	fmt.Println("1:", A1, B1, C1)

	const (
		A2 = iota // 0 : Ξεκινάει στο 0
		B2        // 1 : Αυξάνει κατά 1
		C2        // 2 : Αυξάνει κατά 1
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : Ξεκινάει στο 0 + 1
		B3            // 2 : Αυξάνει κατά 1
		C3            // 3 : Αυξάνει κατά 1
	)

	fmt.Println("3:", A3, B3, C3)

	const (
		Ldate         = 1 << iota //  1 : Μετατόπιση του 1 αριστερά 0 θέσεις.  0000 0001
		Ltime                     //  2 : Μετατόπιση του 1 αριστερά 1 θέσεις.  0000 0010
		Lmicroseconds             //  4 : Μετατόπιση του 1 αριστερά 2 θέσεις.  0000 0100
		Llongfile                 //  8 : Μετατόπιση του 1 αριστερά 3 θέσεις.  0000 1000
		Lshortfile                // 16 : Μετατόπιση του 1 αριστερά 4 θέσεις.  0001 0000
		LUTC                      // 32 : Μετατόπιση του 1 αριστερά 5 θέσεις.  0010 0000
	)

	fmt.Println("Log:", Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC)
}
