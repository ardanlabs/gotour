//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το παράδειγμα προγράμματος παρουσιάζει το πρότυπο καναλιού επικοινωνίας,, περιμένοντας για αποτελέσματα.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForResult()
}

// waitForResult: Σε αυτό το πρότυπο, η αρχική goroutine περιμένει για την παράγωγη
// goroutine να ολοκληρώσει κάποια εργασία ώστε να μεταδώσει το αποτέλεσμα.
func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data"
		fmt.Println("child : sent signal")
	}()

	d := <-ch
	fmt.Println("parent : recv'd signal :", d)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
