//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος επιδεικνύει το πρότυπο καναλιού επικοινωνίας περιορισμένης συγκέντρωση πόρων εργασίας.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	boundedWorkPooling()
}

// boundedWorkPooling: Σε αυτό το πρότυπο, δημιουργείται μια διαθέσιμη συγκέντρωση
// παράγωγων goroutine προκειμένου να εξυπηρετήσουν ένα δεδομένο ποσό εργασίας.
// Η αρχική goroutine προσπελαύνει επαναληπτικά όλες τις εργασίες, ενημερώνοντας
// σχετικά τις διαθέσιμες goroutine. Όταν έχει σηματοδοτηθεί όλος ο όγκος εργασίας,
// τότε το αρχικό κανάλι επικοινωνίας κλείνει, καθαρίζει και οι παράγωγες goroutines
// τερματίζουν την λειτουργία τους.
func boundedWorkPooling() {
	work := []string{"paper", "paper", "paper", "paper", 2000: "paper"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for c := 0; c < g; c++ {
		go func(child int) {
			defer wg.Done()
			for wrk := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, wrk)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	for _, wrk := range work {
		ch <- wrk
	}
	close(ch)
	wg.Wait()

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
