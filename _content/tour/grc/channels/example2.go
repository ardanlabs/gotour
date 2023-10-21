//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει το πρότυπο εξάπλωσης (fan out) ενός καναλιού επικοινωνίας.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanOut()
}

// fanOut: Σε αυτό το πρώτυπο, η αρχική goroutine δημιουργεί 2000 παράγωγες goroutine
// και τις περιμένει να μεταδώσουν τα αποτελέσματά τους.
func fanOut() {
	children := 2000
	ch := make(chan string, children)

	for c := 0; c < children; c++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(c)
	}

	for children > 0 {
		d := <-ch
		children--
		fmt.Println(d)
		fmt.Println("parent : recv'd signal :", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
