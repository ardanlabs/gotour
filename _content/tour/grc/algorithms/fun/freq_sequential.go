//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει τον τρόπο υλοποίησης μιας συνάρτησης,
// που μπορεί να βρει την συχνότητα ενός ρούνου (στμ. rune) σε συγκεκριμένη πρόταση.
package main

import (
	"fmt"
)

func main() {
	sentence := `The quick brown fox jumps over the lazy dog Stay hungry.
	Stay foolish Keep going. Be all in Boldness be my friend Screw it,
	let's do it My life is my message Leave no stone unturned Dream big.
	Pray bigger`

	print(sequential(sentence))
}

func sequential(text string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range text {
		m[r]++
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
