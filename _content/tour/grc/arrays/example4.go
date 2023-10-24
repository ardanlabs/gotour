//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ότι η έκφραση for range
// έχει τόσο σημειολογία τιμής, όσο και σημειολογία δείκτη διεύθυνσης.
package main

import "fmt"

func main() {

	// Χρησιμοποιώντας την μορφή σημειολογίας δείκτη διεύθυνσης,
	// της έκφρασης for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1])
		}
	}

	// Χρησιμοποιώντας την μορφή σημειολογίας τιμής, της έκφρασης
	// for range.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}

	// Χρησιμοποιώντας την μορφή σημειολογίας τιμής της έκφρασης for range,
	// όμως με πρόσβαση σημειολογίας δείκτη διεύθυνσης.
	// ΜΗΝ ΤΟ ΚΑΝΕΤΕ ΑΥΤΟ.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	for i, v := range &friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	}
}
