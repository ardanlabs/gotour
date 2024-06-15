//go:build OMIT || norun

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως η έκφραση for range
// έχει σημειολογία, τόσο τιμής όσο και δείκτη διεύθυνσης.
package main

import "fmt"

func main() {

	// Χρησιμοποιώντας την μορφή της έκφρασης for range,
	// που έχει σημειολογία τιμής.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}

	// Χρησιμοποιώντας την μορφή της έκφρασης for range,
	// που έχει σημειολογία δείκτη διεύθυνσης.
	friends = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
}
