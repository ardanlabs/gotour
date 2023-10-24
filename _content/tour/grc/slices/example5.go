//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί πως κανείς πρέπει να είναι
// προσεχτικός όταν προσθέτει σε μια φέτα, όταν υπάρχει αναφορά σε κάποιο στοιχείο.
package main

import "fmt"

type user struct {
	likes int
}

func main() {

	// Δηλώστε μια φέτα με 3 user.
	users := make([]user, 3)

	// Μοιραστείτε τον χρήστη στον δείκτη 1.
	shareUser := &users[1]

	// Προσθέστε ένα like για τον χρήστη, που μοιράστηκε προηγουμένως.
	shareUser.likes++

	// Παρουσιάστε τον αριθμό των like, για όλους τους χρήστες.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Προσθέστε έναν νέο χρήστη.
	users = append(users, user{})

	// Προσθέστε ακόμα ένα like για τον χρήστη, που μοιράστηκε προηγουμένως.
	shareUser.likes++

	// Παρουσιάστε τον αριθμό των like, για όλους τους χρήστες.
	fmt.Println("*************************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Παρατηρείστε ότι το τελευταίο like δεν έχει καταγραφεί.
}
