//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος συνδυασμού περιορισμών τύπου
// που βασίζονται στον τύπο και στην συμπεριφορά.
package main

import "fmt"

// Ορίζοντας δύο πραγματικούς τύπους που υλοποιούν μια μέθοδο τύπου με το όνομα match.

type person struct {
	name  string
	email string
}

func (p person) match(v person) bool {
	return p.name == v.name
}

type food struct {
	name     string
	category string
}

func (f food) match(v food) bool {
	return f.name == v.name
}

// Η διεπαφή matcher ορίζει δύο περιορισμούς. Πρώτα, περιορίζει τα δεδομένα
// στο τύπο που είναι αποδεκτός. Δεύτερον, περιορίζει την συμπεριφορά των δεδομένων.
// Η μέθοδος τύπου match απαιτεί ότι μια τιμή τύπου T (που θα προσδιοριστεί αργότερα)
// θα είναι η είσοδος της μεθόδου τύπου.

// Σημείωση: Ο κατάλογος των τύπων στην διεπαφή δεν είναι απαραίτητος προκειμένου
//			 να επιτύχει το ταίριασμα. Αυτό που προσπαθούμε να παρουσιάσουμε είναι
// 			 είναι ο τρόπος που μπορούν να συνδυαστούν ο κατάλογος τύπων και η συμπεριφορά.

type matcher[T any] interface {
	person | food
	match(v T) bool
}

// Η συνάρτηση match δηλώνει ότι η τιμή τύπου T πρέπει να υλοποιεί την διεπαφή
// matcher και χρησιμοποιείται για τα ορίσματα φέτας και τιμής της συνάρτησης.

func match[T matcher[T]](list []T, find T) int {
	for i, v := range list {
		if v.match(find) {
			return i
		}
	}
	return -1
}

// =============================================================================

func main() {
	people := []person{
		{name: "bill", email: "bill@email.com"},
		{name: "jill", email: "jill@email.com"},
		{name: "tony", email: "tony@email.com"},
	}
	findPerson := person{name: "tony"}

	i := match(people, findPerson)
	fmt.Printf("Match: Idx: %d for %s\n", i, findPerson.name)

	foods := []food{
		{name: "apple", category: "fruit"},
		{name: "carrot", category: "veg"},
		{name: "chicken", category: "meat"},
	}
	findFood := food{name: "apple"}

	i = match(foods, findFood)
	fmt.Printf("Match: Idx: %d for %s\n", i, findFood.name)
}
