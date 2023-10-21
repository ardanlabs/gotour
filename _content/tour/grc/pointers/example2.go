//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί η βασική έννοια της χρήσης δείκτη διεύθυνσης
// προκειμένου να μοιραστεί κανείς δεδομένα.
package main

func main() {

	// Δηλώστε μεταβλητή τύπου int με τιμή ίση με 10.
	count := 10

	// Παρουσιάστε την "τιμή της" και την "διεύθυνση μνήμης της" μεταβλητής count.
	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")

	// Περάστε στην συνάρτηση την "διεύθυνση της" μεταβλητής count.
	increment(&count)

	println("count:\tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

// Η increment δηλώνει την count ως μεταβλή δείκτη διεύθυνσης της οποίας η τιμή είναι
// πάντα μια διεύθυσνη μνήμης και δείχνει σε τιμές τύπου int.
//
//go:noinline
func increment(inc *int) {

	// Αυξήστε την "τιμή της" μεταβλητής στην οποία "ο δείκτης διεύθυνσης δείχνει προς".
	*inc++

	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}
