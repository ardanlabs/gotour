//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστούν τα συστατικά μιας φέτας. Η φέτα έχει
// μήκος, χωρητικότητα καθώς και τον υποκείμενο πίνακα.
package main

import "fmt"

func main() {

	// Δημιουργείστε μια φέτα με μήκος 5 στοιχεία και χωρητικότητα 8.
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	inspectSlice(fruits)
}

// Η inspectSlice αποκαλύπτει την επικεφαλίδα τηε φέτας, προς επισκόπηση.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}
