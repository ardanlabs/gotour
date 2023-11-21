//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος πραγματικών
// υλοποιήσεων συναρτήσεων εκτύπωσης, που μπορούν να λειτουργήσουν μόνο
// με φέτες συγκεκριμένου τύπου.
package main

import (
	"fmt"
)

func printNumbers(numbers []int) {
	fmt.Print("Numbers: ")

	for _, num := range numbers {
		fmt.Print(num, " ")
	}

	fmt.Print("\n")
}

func printStrings(strings []string) {
	fmt.Print("Strings: ")

	for _, str := range strings {
		fmt.Print(str, " ")
	}

	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3}
	printNumbers(numbers)

	strings := []string{"A", "B", "C"}
	printStrings(strings)
}
