//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος επιδεικνύει το πρότυπο καναλιού επικοινωνίας
// συγκέντρωσης πόρων.
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	pooling()
}

// Πρότυπο συγκέντρωση πόρων (στμ. pooling): Σε αυτό το πρότυπο, η αρχική
// ρουτίνα συνεκτέλεσης της Go σηματοδοτεί την ύπαρξη 100 τεμαχίων εργασίας
// στο διαθέσιμο απόθεμα παράγωγων ρουτίνων συνεκτέλεσης της Go, οι οποίες
// περιμένουν να πραγματοποιήσουν εργασία.
func pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : recv'd signal : %s\n", child, d)
			}
			fmt.Printf("child %d : recv'd shutdown signal\n", child)
		}(c)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "data"
		fmt.Println("parent : sent signal :", w)
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
