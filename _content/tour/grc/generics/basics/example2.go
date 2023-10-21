//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος συγγραφής μιας συνάρτησης
// που παρέχει μια λύση άδειας διεπαφής, η οποία χρησιμοποιεί διαβεβαιώσεις τύπων για τις
// διαφορετικές πραγματικές φέτες που μπορεί να υποστηρίξει. Βασικά μεταφέραμε τις
// παραπάνω συναρτήσεις σε δηλώσεις περιπτώσεων case.
package main

import (
	"fmt"
)

func printAssert(v interface{}) {
	fmt.Print("Assert: ")

	switch list := v.(type) {
	case []int:
		for _, num := range list {
			fmt.Print(num, " ")
		}

	case []string:
		for _, str := range list {
			fmt.Print(str, " ")
		}
	}

	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3}
	printAssert(numbers)

	strings := []string{"A", "B", "C"}
	printAssert(strings)
}
