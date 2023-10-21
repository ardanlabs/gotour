//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκιεμένου να παρουσιαστεί πως μεταβλητές ενός ανώνυμου τύπου μπορούν
// να εκχωρηθούν σε μεταβλητές ενός επώνυμου τύπου, όταν είναι
// ίδιοι.
package main

import "fmt"

// Ο example αναπαριστά έναν τύπο με διαφορετικά πεδία.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Δηλώστε μια μεταβλητή ενός ανώνυμου τύπου και δώστε της αρχική τιμή
	// χρησιμοποιώντας μια ρητή κατασκευή struct.
	e := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Δημιουργείστε μια τιμή τύπου example.
	var ex example

	// Εκχωρείστε την τιμή του ανώνυμου τύπου struct
	// στην τιμή του επώνυμου τύπου struct.
	ex = e

	// Παρουσιάστε τις τιμές.
	fmt.Printf("%+v\n", ex)
	fmt.Printf("%+v\n", e)
	fmt.Println("Flag", e.flag)
	fmt.Println("Counter", e.counter)
	fmt.Println("Pi", e.pi)
}
