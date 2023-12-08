//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ο τρόπος συγγραφής μιας
// συνάρτησης, που παρέχει μια λύση αντανάκλασης, η οποία επιτρέπει να παραληφθεί
// και να τυπωθεί μια φέτα κάθε τύπου. Αυτή είναι μια συνάρτηση γενικού
// προγραμματισμού, εξαιτίας του πακέτου reflect.
package main

import (
	"fmt"
	"reflect"
)

func printReflect(v interface{}) {
	fmt.Print("Reflect: ")

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < val.Len(); i++ {
		fmt.Print(val.Index(i).Interface(), " ")
	}

	fmt.Print("\n")
}

func main() {
	numbers := []int{1, 2, 3}
	printReflect(numbers)
	print(numbers)

	strings := []string{"A", "B", "C"}
	printReflect(strings)
}
