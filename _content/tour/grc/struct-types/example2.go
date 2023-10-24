//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος δήλωσης και εκχώρησης
// αρχικής τιμής ανώνυμων τύπων struct.
package main

import "fmt"

func main() {

	// Δηλώστε μια μεταβλητή ενός ανώνυμου τύπου που λαμβάνει την
	// μηδενική τιμή του.
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Παρουσιάστε την τιμή.
	fmt.Printf("%+v\n", e1)

	// Δηλώστε μια μεταβλητή ενός ανώνυμου τύπου και δώστε αρχική τιμή,
	// χρησιμοποιώντας μια ρητή κατασκευή struct.
	e2 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Παρουσιάστε τις τιμές.
	fmt.Printf("%+v\n", e2)
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
