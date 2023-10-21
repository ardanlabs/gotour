//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος δήλωσης και εκχώρησης αρχικής τιμής σε τύπους struct.
package main

import "fmt"

// Ο example αναπαριστά ένα τύπο με διαφορετικά πεδία.
type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	// Δηλώστε μια μεταβλητή τύπου example και εχκωρείστε της, την
	// μηδενική της τιμή.
	var e1 example

	// ΠΑρουσιάστε την τιμή.
	fmt.Printf("%+v\n", e1)

	// Δηλώστε μια μεταβλητή τύπου example και δώστε της αρχική τιμή κάνοντας χρήση
	// ενός ρητού struct.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Παρουσιάστε τις τιμές των πεδίων.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
