//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο συγγραφής μιας ταξινόμηση σωρού.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(10)
	fmt.Println("Before:", numbers)

	heapSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func heapSort(numbers []int) []int {

	// Χωρείστε την λίστα στην μέση και επεξεργαστείτε το μπροστινό μισό
	// της λίστας, μετακινώντας την μεγαλύτερη τιμή που συναντάτε μπροστά
	// και στην συνέχεια κάντε το ίδιο για την αμέσως μικρότερη.

	for index := (len(numbers) / 2) - 1; index >= 0; index-- {
		numbers = moveLargest(numbers, len(numbers), index)
	}

	// Πάρτε την λίστα και αρχίστε να μετακινείτε αριθμούς προς μια νέα
	// ταξινομημένη λίστα. Πάρτε τον αριθμό στην πρώτη θέση και απομακρύνετε
	// τον προς την νέα λίστα η οποία θα περιέχει την τελική ταξινόμηση.
	// Στην συνέχεια, μεταφέρετε τον μεγαλύτερο αριθμό που βρίσκετε ξανά
	// μπροστά στην λίστα.

	size := len(numbers)
	for index := size - 1; index >= 1; index-- {
		numbers[0], numbers[index] = numbers[index], numbers[0]
		size--
		numbers = moveLargest(numbers, size, 0)
	}

	return numbers
}

// Η moveLargest ξεκινάει στους δείκτες θέσης που προσδιορίζονται στην λίστα
// και προσπαθεί να μετακινήσει τον μεγαλύτερο αριθμό που μπορεί να βρει
// σε αυτή την θέση στην λίστα.
func moveLargest(numbers []int, size int, index int) []int {

	// Υπολογείστε την απόκλιση του δείκτη ώστε οι αριθμοί στην λίστα να
	// μπορούν να συγκριθούν και να αλλάξουν θέση αν χρειαστεί.
	// index 0: cmpIdx1: 1 cmpIdx2:  2   index 5: cmpIdx1: 11 cmpIdx2: 12
	// index 1: cmpIdx1: 3 cmpIdx2:  4   index 6: cmpIdx1: 13 cmpIdx2: 14
	// index 2: cmpIdx1: 5 cmpIdx2:  6   index 7: cmpIdx1: 15 cmpIdx2: 16
	// index 3: cmpIdx1: 7 cmpIdx2:  8   index 8: cmpIdx1: 17 cmpIdx2: 19
	// index 4: cmpIdx1: 9 cmpIdx2: 10   index 9: cmpIdx1: 19 cmpIdx2: 20
	cmpIdx1, cmpIdx2 := 2*index+1, 2*index+2

	// Αποθηκεύστε τον προσδιορίσμένο δείκτη ως τον δείκτη με την τρέχουσα μεγαλύτερη τιμή.
	largestValueIdx := index

	// Ελέγξτε αν η τιμή στον πρώτη δείκτη απόκλισης είναι μεγαλύτερη
	// από την τιμή στον τρέχοντα δείκτη μεγαλύτερης τιμής. Αν είναι έτσι
	// αποθηκεύστε αυτή την θέση δείκτη.
	if cmpIdx1 < size && numbers[cmpIdx1] > numbers[largestValueIdx] {
		largestValueIdx = cmpIdx1
	}

	// Ελέγξτε αν ο δεύτερος δείκτης απόκλισης είναι εντός ορίων και
	// ότι είναι μεγαλύτερος από την τιμή στον τρέχοντα μεγαλύτερο δείκτη.
	// Αν αυτό ισχύει, αποθηκεύστε την θέση δείκτη.
	if cmpIdx2 < size && numbers[cmpIdx2] > numbers[largestValueIdx] {
		largestValueIdx = cmpIdx2
	}

	// Αν βρείτε μια μεγαλύτερη τιμή από αυτή που υπάρχει στο συγκεκριμένο δείκτη,
	// τότε ανταλλάξτε αυτούς τους αριθμούς και μετά επαναλάβετε την αναζήτηση για
	// περισσότερους αριθμούς προς ανταλλαγή από αυτό το σημείο στη λίστα.
	if largestValueIdx != index {
		numbers[index], numbers[largestValueIdx] = numbers[largestValueIdx], numbers[index]
		numbers = moveLargest(numbers, size, largestValueIdx)
	}

	return numbers
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
