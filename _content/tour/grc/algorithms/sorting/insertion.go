//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής μιας ταξινόμησης εισαγωγής.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(10)
	fmt.Println("Before:", numbers)

	insertionSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func insertionSort(numbers []int) {
	var n = len(numbers)

	// Διατρέξτε τους αριθμούς από αριστερά προς τα δεξιά.
	// Μέσα από κάθε επανάληψη του εξωτερικού βρόγχου μεταφέρουμε
	// τιμές από δεξιά στα αριστερά εντός του πίνακα όταν είναι
	// μεγαλύτερες από την τιμή που προηγείται.

	for i := 1; i < n; i++ {
		j := i

		// Για τον δεδομένο αρχικό δείκτη θέσης i, αναζητείστε
		// μικρότερες τιμές προκειμένου να τις μεταφέρετε αριστερά
		// στον κατάλογο των αριθμών.

		for j > 0 {

			// Είναι η τιμή αριστερά μεγαλύτερη από την
			// δεξιά; Αν είναι αληθές, εναλλάξτε τις δύο
			// τιμές.

			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}

			// Επισκεφθείτε το στοιχείο από δεξιά προς τα αριστερά.

			j--
		}
	}
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
