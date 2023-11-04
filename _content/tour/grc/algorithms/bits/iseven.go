//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο ελέγχου
// ενός ακεραίου ως ζυγού ή μονού, χρησιμοποιώντας λειτουργίες
// για bit.
package main

import (
	"fmt"
)

func main() {

	fmt.Println(8, ":", IsEven(8))
	fmt.Println(15, ":", IsEven(15))
	fmt.Println(4, ":", IsEven(4))
}

// Η IsEven ελέγχει αν ένας ακέραιος είναι ζυγός.
func IsEven(num int) bool {

	// χρησιμοποιήστε τον τελεστή AND για bit, προκειμένου να εξακριβώσετε
	// αν το λιγότερο σημαντικό bit (στμ. LSB) είναι 0.

	// Χρήσιμη πηγή: https://catonmat.net/low-level-bit-hacks
	// 0 & 1 = 0 (ζυγός αριθμός)
	// 1 & 1 = 1 (μονός αριθμός)

	return num&1 == 0
}
