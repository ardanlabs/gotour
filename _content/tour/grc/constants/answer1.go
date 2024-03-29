//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δηλώστε μια σταθερά χωρίς τύπο και μια με τύπο και παρουσιάστε τις τιμές τους.
//
// Πολλαπλασιάστε δυο ρητές σταθερές σε μια μεταβλητή με τύπο και παρουσιάστε την τιμή.
package main

import "fmt"

const (
	// Η server είναι η διεύθυνση IP προς σύνδεση.
	server = "124.53.24.123"

	// Η port είναι η θύρα που θα χρησιμοποιηθεί, προκειμένου να γίνει η σύνδεση.
	port int16 = 9000
)

func main() {

	// Παρουσιάστε την πληροφορία για την server.
	fmt.Println(server)
	fmt.Println(port)

	// Υπολογίστε τον αριθμό των λεπτών σε 5320 δευτερόλεπτα.
	minutes := 5320 / 60.0
	fmt.Println(minutes)
}
