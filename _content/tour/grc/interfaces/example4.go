//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως δεν είναι πάντα δυνατόν να πάρει κανείς την διεύθυνση μιας τιμής.
package main

import "fmt"

// Ο duration είναι ένας επώνυος τύπος με βασικό τύπο έναν ακέραιο.
type duration int

// Η notify υλοποιεί την διεπαφή notifier.
func (d *duration) notify() {
	fmt.Println("Sending Notification in", *d)
}

func main() {
	duration(42).notify()

	// ./example3.go:18: cannot call pointer method on duration(42)
	// ./example3.go:18: cannot take the address of duration(42)
}
