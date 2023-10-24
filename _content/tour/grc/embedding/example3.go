//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως λειτουργούν οι
// ενσωματωμένοι τύποι, με διεπαφές.
package main

import "fmt"

// Ο notifier είναι μια διεπαφή, που ορίζει συμπεριφορά ειδοποιήσεων.
// type behavior.
type notifier interface {
	notify()
}

// Ο user ορίζει έναν χρήστη στο πρόγραμμα.
type user struct {
	name  string
	email string
}

// Η notify υλοποιεί μια μέθοδο τύπου, που ειδοποιεί τους
// χρήστες για διαφορετικά γεγονότα.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// Ο admin αναπαριστά έναν χρήστη διαχειριστή, με προνόμια.
type admin struct {
	user
	level string
}

func main() {

	// Δημιουργήστε έναν χρήστη admin.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Στείλτε στον χρήστη admin μια ειδοποίηση.
	// Η υλοποίηση της διεπαφής του ενσωματωμένου εσωτερικού
	// τύπου "προωθείται" στον εξωτερικό τύπο.
	sendNotification(&ad)
}

// Η sendNotification δέχεται τιμές που υλοποιούν την διεπαφή
// notifier και αποστέλλει ειδοποιήσεις.
func sendNotification(n notifier) {
	n.notify()
}
