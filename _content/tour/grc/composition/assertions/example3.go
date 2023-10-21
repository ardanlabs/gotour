//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί πως τα σύνολα μεθόδων τύπων μπορούν να επηρεάσουν την συμπεριφορά.
package main

import "fmt"

// Ο user ορίζει έναν χρήστη στο σύστημα.
type user struct {
	name  string
	email string
}

// Η String υλοποιεί την διεπαφή fmt.Stringer.
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {

	// Δημιουργείστε μια τιμή τύπου user.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Παρουσιάστει τις τιμές.
	fmt.Println(u)
	fmt.Println(&u)
}
