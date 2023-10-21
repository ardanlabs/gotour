//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος επιδεικνύει το πρότυπο καναλιού επικοινωνίας απόρριψης εργασίας.
// This sample program demonstrates the drop channel pattern.
package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

// Απόρριψη εργασίας (drop): Σε αυτό το πρότυπο, η αρχική goroutine σηματοδοτεί
// την παρουσία 2000 τεμαχίων εργασίας σε μια παράγωγη goroutine η οποία δεν είναι
// σε θέση να χειριστεί όλο το έργο. Αν η αρχική goroutine πραγματοποιεί μια
// αποστολή και η παράγωγη goroutine δεν είναι σε κατάσταση ετοιμότητας να παραλάβει
// αυτή η εργασία απορρίπτεται και πετιέται.
func drop() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : recv'd signal :", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", w)
		default:
			fmt.Println("parent : dropped data :", w)
		}
	}

	close(ch)
	fmt.Println("parent : sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
