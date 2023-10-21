//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο ελέγχου αν ένας ακέραιος
// είναι ένα παλίνδρομο ή όχι.
package main

import "fmt"

func main() {
	tt := []int{-1, 1, 9, 10, 1001, 125}

	for _, input := range tt {
		success := IsPalindrome(input)

		switch success {
		case true:
			fmt.Printf("%d is a palindrome\n", input)

		case false:
			fmt.Printf("%d is NOT a palindrome\n", input)
		}
	}
}

// Η IsPalindrome ελέγχει αν ένας ακέραιος είναι Παλίνδρομο.
func IsPalindrome(input int) bool {

	// Ένας αρνητικός ακέραιος δεν μπορεί να είναι παλίνδρομος.
	if input < 0 {
		return false
	}

	// Ένας ακέραιος με μήκος μόνο ενός αριθμού είναι ένα παλίνδρομο.
	if input >= 0 && input < 10 {
		return true
	}

	// Αντιστρέψτε τα ψηφία στον ακέραιο.
	rev := Reverse(input)

	return input == rev
}

// Η Reverse παίρνει τον συγκεκριμένο ακέραιο και τον αντιστρέφει.
func Reverse(num int) int {

	// Δημιουργήστε την result με την μηδενική της τιμή.
	var result int

	// Επαναλάβετε, εως ότου η num να ισούται με το μηδέν.
	for num != 0 {

		// Πραγματοποιείστε την λειτουργία υπολοίπου ακέραιας διαίρεσης (modulus)
		// προκειμένου να πάρετε το τελευταίο ψηφίο από την τιμή που έχει ανατεθεί στην num.
		// https://www.geeksforgeeks.org/find-first-last-digits-number/
		// Πχ. Για num = 125, τελευταίο = 5
		last := num % 10

		// Πολλαπλασιάστε την τρέχουσα τιμή
		// της result με 10 προκειμένου να
		// μεταθέσετε τα ψηφία στην τρέχουσα
		// τιμή της result προς τα αριστερά.
		// Πχ. Για result = 5, result = 50
		result = result * 10

		// Προσθέστε το ψηφίο που πήραμε από το τέλος της
		// num στην result.
		// Πχ. Για result = 21 και last = 5, result = 215
		result += last

		// Απομακρύνετε το ψηφίο που μόλις αντιστρέψαμε από την num.
		// Πχ. Για num = 125, num = 12
		num = num / 10
	}

	return result
}
