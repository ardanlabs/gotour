//go:build OMIT || nobuild

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί το γεγονός, ότι πίνακες
// διαφορετικών μεγεθών ανήκουν σε διαφορετικούς τύπους.
package main

import "fmt"

func main() {

	// Δηλώστε έναν πίνακα με 5 ακέραιους, ο οποίος λαμβάνει την μηδενική
	// του τιμή.
	var five [5]int

	// Δηλώστε έναν πίνακα με 4 ακέραιους, ο οποίος λαμβάνει ως αρχική τιμή,
	//  κάποιες τιμές.
	four := [4]int{10, 20, 30, 40}

	// Εκχωρείστε τον ένα πίνακα, στον άλλο.
	five = four

	// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment

	fmt.Println(four)
	fmt.Println(five)
}
