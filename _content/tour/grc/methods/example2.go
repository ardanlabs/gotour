//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως δηλώνονται μέθοδοι τύπου
// για επώνυμο τύπο.
package main

import "fmt"

// Η duration είναι ένας επώνυμος τύπος που αναπαριστά μια διάρκεια
// χρόνου σε Nanosecond.
type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

// Η setHours θέτει την συγκεκριμένη τιμή ωρών.
func (d *duration) setHours(h float64) {
	*d = duration(h) * hour
}

// Η hours επιστρέφει την διάρκεια ως αριθμό κινητής υποδιαστολής ωρών.
func (d duration) hours() float64 {
	hour := d / hour
	nsec := d % hour
	return float64(hour) + float64(nsec)*(1e-9/60/60)
}

func main() {

	// Δηλώστε μια μεταβλητή τύπου duration που λαμβάνει την μηδενική τιμή του τύπου.
	var dur duration

	// Αλλάξτε την τιμή της dur προκειμένου να ισούται πέντε ώρες.
	dur.setHours(5)

	// Παρουσιάστε την νέα τιμή της dur.
	fmt.Println("Hours:", dur.hours())
}
