//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, όπου γίνεται εξερεύνηση του τρόπου λειτουργίας των αναθέσεων
// σε διεπαφές, όταν αποθηκεύονται τιμές μέσα στην διεπαφή.
package main

import (
	"fmt"
	"unsafe"
)

// Ο notifier παρέχει υποστήριξη για γεγονότα ειδοποίησης.
type notifier interface {
	notify()
}

// Ο user αναπαριστά έναν χρήστη στο σύστημα.
type user struct {
	name string
}

// Η notify υλοποιεί την διεπαφή notifier.
func (u user) notify() {
	fmt.Println("Alert", u.name)
}

func inspect(n *notifier, u *user) {
	word := uintptr(unsafe.Pointer(n)) + uintptr(unsafe.Sizeof(&u))
	value := (**user)(unsafe.Pointer(word))
	fmt.Printf("Addr User: %p  Word Value: %p  Ptr Value: %v\n", u, *value, **value)
}

func main() {

	// Δημιουργείστε μια διεπαφή notifier και μια πραγματική τιμή τύπου.
	var n1 notifier
	u := user{"bill"}

	// Αποθηκεύστε ένα αντίγραφο της τιμής user στην τιμή διεπαφής notifier.
	n1 = u

	// Παρατηρούμε ότι η διεπαφή έχει το δικό της αντίγραφο.
	// Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n1, &u)

	// Κάντε ένα αντίγραφο της τιμής διεπαφής.
	n2 := n1

	// Παρατηρούμε ότι η διεπαφή μοιράζεται την ίδια τιμή που είναι αποθηκευμένη
	// στην τιμή διεπαφής n1.
	// Addr User: 0x1040a120  Word Value: 0x10427f70  Ptr Value: {bill}
	inspect(&n2, &u)

	// Αποθηκεύστε ένα αντίγραφο της τιμής διεύθυνσης μνήμης του χρήστη
	// στην τιμή της διεπαφής notifier.
	n1 = &u

	// Παρατηρούμε ότι η διεπαφή μοιράζεται την τιμή της μεταβλητής u, άμεσα.
	// Δεν υπάρχει αντίγραφο.
	// Addr User: 0x1040a120  Word Value: 0x1040a120  Ptr Value: {bill}
	inspect(&n1, &u)
}
