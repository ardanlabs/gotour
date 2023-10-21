//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως μεγεθύνεται μια φέτα χρησιμοποιώντας
// την προεγκατεστημένη συνάρτηση append και πως η append μεγεθύνει την χωρητικότητα του
// υποκείμενου πίνακα.
package main

import "fmt"

func main() {

	// Δηλώστε μια φέτα συμβολοσειρών με τιμή nil.
	var data []string

	// Κρατήστε την χωρητικότητα της φέτας.
	lastCap := cap(data)

	// Προσθέστε ~100k συμβολοσειρές στην φέτα.
	for record := 1; record <= 1e5; record++ {

		// Χρησιμοποιείστε τη προεγκατεστημένη συνάρτηση append προκειμένου να προσθέσετε στην φέτα.
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// Όταν η χωρτηικότητα της φέτας αλλάζει, παρουσιάστε τις αλλαγές.
		if lastCap != cap(data) {

			// Υπολογίστε το ποσοστό της αλλαγής.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Αποθηκεύστε την νέα τιμή της χωρητικότητας.
			lastCap = cap(data)

			// Παρουσιάστε τα αποτελέσματα.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}
