// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος ταυτόχρονης εκτέλεσης παρουσιάζει τον τρόπο
// υλοποίησης μιας συνάρτησης, που μπορεί να βρει την συχνότητα, με την οποία
// χρησιμοποιείται ένας ρούνος (στμ. rune) σε μια συγκεκριμένη πρόταση.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	sentence := `The quick brown fox jumps over the lazy dog Stay hungry.
	Stay foolish Keep going. Be all in Boldness be my friend Screw it,
	let's do it My life is my message Leave no stone unturned Dream big.
	Pray bigger`

	print(concurrent(sentence))
}

func concurrent(text string) map[rune]int {
	m := make(map[rune]int)    // Πίνακας κατακερματισμού με τελικό αποτέλεσμα.
	g := runtime.GOMAXPROCS(0) // Αριθμός των goroutine.
	l := len(text)             // Αριθμός των byte για επεξεργασία.
	b := l / g                 // Αριθμός θέσεων αποθήκευσης, μία ανά goroutine

	// Παραλαμβάνει το αποτέλεσμα επεξεργασίας κάθε θέσης αποθήκευσης, από
	// μια ρουτίνα συνεκτέλεσης της Go.
	ch := make(chan map[rune]int, g)

	// Δημιουργήστε πλήθος g goroutine.

	for i := 0; i < g; i++ {
		str := i * b   // Αρχική θέση δείκτη του χώρου αποθήκευσης.
		end := str + b // Τελική θέση δείκτη του χώρου αποθήκευσης.
		if i == g-1 {  // Η τελευταία θέση αποθήκευσης παίρνει τα υπολειπόμενα byte.
			end = end + (l - end)
		}

		go func() {
			m := make(map[rune]int)

			defer func() {
				ch <- m
			}()

			// Αυτή η G επεξεργάζεται τον χώρο αποθήκευσης, σειριακά.
			for _, r := range text[str:end] {
				m[r]++
			}
		}()
	}

	// Περιμένετε την επιστροφή των αποτελεσμάτων κάθε χώρου αποθήκευσης
	// και επεξεργαστείτε τα στον τελικό σχεσιακό πίνακα.

	for i := 0; i < g; i++ {
		result := <-ch
		for rk, rv := range result {
			m[rk] = m[rk] + rv
		}
	}

	return m
}

func print(m map[rune]int) {
	var cols int

	for r := 65; r < 65+26; r++ {
		v := m[rune(r)]
		fmt.Printf("%q:%d, ", rune(r), v)

		v = m[rune(r+32)]
		fmt.Printf("%q:%d, ", rune(r+32), v)

		cols++
		if cols == 5 {
			fmt.Print("\n")
			cols = 0
		}
	}
}
