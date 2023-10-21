//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// ΔΕίγμα προγράμματος προκειμένου να παρουσιαστεί τι συμβαίνει όταν ο εξώτερος και ο εσώτερος τύπος
// υλοποιούν την ίδια διεπαφή.
package main

import "fmt"

// Ο notifier είναι μια διεπαφή που ορίζει συμπεριφορά ειδοποιήσεων.
type notifier interface {
	notify()
}

// Ο user ορίζει έναν χρήστη στο πρόγραμμα.
type user struct {
	name  string
	email string
}

// Η notify υλοποιεί μια μέθοδο που ενημερώνει τους χρήστες
// για διαφορετικά γεγονότα.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

// Ο admin αναπαριστά έναν χρήστη διαχειριστή με προνόμια.
type admin struct {
	user
	level string
}

// Η notify υλοποιεί μια μέθοδο τύπου που ειδοποιεί τους διαχειριστές
// για διάφορα γεγονότα.
func (a *admin) notify() {
	fmt.Printf("Sending admin Email To %s<%s>\n",
		a.name,
		a.email)
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
	// Η υλοποίηση της διεπαφής του ενσωματωμένου εσωτερικού τύπου
	// ΔΕΝ "προωθείται" στον εξωτερικό τύπο.
	sendNotification(&ad)

	// Μπορούμε να έχουμε πρόσβαση στην μέθοδο τύπου του εσωτερικού τύπου απευθείας.
	ad.user.notify()

	// Η μέθοδος του εσωτερικού τύπου ΔΕΝ προωθείται.
	ad.notify()
}

// Η sendNotification αποδέχεται τιμές που υλοποιούν την διεπαφή notifier
// interface και στέλνει ειδοποιήσεις.
func sendNotification(n notifier) {
	n.notify()
}
