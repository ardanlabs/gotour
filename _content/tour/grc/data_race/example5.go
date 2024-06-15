//go:build OMIT || norun

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, προκειμένου να παρουσιαστεί ότι οι σχεσιακοί πίνακες
// δεν είναι ασφαλείς για χρήση ταυτόχρονης εκτέλεσης, από κατασκευής.
// Το περιβάλλον εκτέλεσης θα ανιχνεύσει εγγραφες ταυτόχρονης εκτέλεσης
// και θα προκαλέσει κατάσταση panic.
package main

import (
	"fmt"
	"sync"
)

// Η scores κρατάει τιμές, που αυξάνονται από πολλές ρουτίνες ]
// συνεκτέλεσης της Go.
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final scores:", scores)
}
