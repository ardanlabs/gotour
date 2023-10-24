//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος επιδεικνύει το πρότυπο καναλιού επικοινωνίας
// διασκορπισμού (στμ. fan-out) με σηματοφόρο (στμ. semaphore).
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fanOutSem()
}

// fanOutSem: Σε αυτό το πρότυπο, ένας σηματοφόρος προστίθεται στο πρότυπο
// διασκορπισμού, προκειμένου να περιοριστεί ο αριθμός των παράγωγων
// ρουτίνων συνεκτέλεσης της Go, που είναι δυνατόν να χρονοδρομολογηθούν προς
// εκτέλεση.
func fanOutSem() {
	children := 2000
	ch := make(chan string, children)

	g := runtime.GOMAXPROCS(0)
	sem := make(chan bool, g)

	for c := 0; c < children; c++ {
		go func(child int) {
			sem <- true
			{
				t := time.Duration(rand.Intn(200)) * time.Millisecond
				time.Sleep(t)
				ch <- "data"
				fmt.Println("child : sent signal :", child)
			}
			<-sem
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
