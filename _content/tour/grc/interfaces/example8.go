//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί το συντακτικό και ο μηχανισμός
// των εναλλαγών τύπων και της άδειας διεπαφής.
package main

import "fmt"

func main() {

	// Η fmt.Println μπορεί να κληθεί με τιμές κάθε τύπου.
	fmt.Println("Hello, world")
	fmt.Println(12345)
	fmt.Println(3.14159)
	fmt.Println(true)

	// Πως μπορούμε να επιτύχουμε το ίδιο αποτέλεσμα;
	myPrintln("Hello, world")
	myPrintln(12345)
	myPrintln(3.14159)
	myPrintln(true)

	// - Μια διεπαφή ικανοποιείται από ένα κομμάτι δεδομένων όταν αυτά τα
	//   δεδομένα επιδεικνύουν το πλήρες σύνολο συμπεριφορών μεθόδων τύπου,
	//   που ορίζονται από την διεπαφή.
	// - Η άδεια διεπαφή δεν ορίζει κανένα σύνολο συμπεριφορών μεθόδων
	//   τύπου και επομένως δεν απαιτεί καμία μέθοδο τύπου, από τα δεδομένα
	//   που αποθηκεύει.

	// - Η άδεια διεπαφή δεν λέει τίποτα σχετικά με τα δεδομένα, που είναι
	//   αποθηκευμένα στην διεπαφή.
	// - Θα πρέπει να πραγματοποιηθούν έλεγχοι κατά το στάδιο της εκτέλεσης,
	//   προκειμένου να μάθει κανείς οτιδήποτε για τα δεδομένα, που είναι
	//   αποθηκευμένα στην άδεια διεπαφή.
	// - Πραγματοποιήστε αποσύνδεση γύρω από καλά ορισμένες συμπεριφορές και
	//   χρησιμοποιήστε την άδεια διεπαφή σαν μια εξαίρεση, όταν αυτό είναι
	//   λογικό και πρακτικό.
}

func myPrintln(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("Is string  : type(%T) : value(%s)\n", v, v)
	case int:
		fmt.Printf("Is int     : type(%T) : value(%d)\n", v, v)
	case float64:
		fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
	default:
		fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
	}
}
