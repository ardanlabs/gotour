//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Αυτό το δείγμα προγράμματος παρουσιάζει το πρότυπο καναλιού επικοινωνίας
// αναμονής για εργασία.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	waitForTask()
}

// waitForTask: Σε αυτό το πρότυπο, η αρχική ρουτίνα συνεκτέλεσης της Go
// αποστέλλει ένα σήμα σε μια παράγωγη ρουτίνα συνεκτέλεσης της Go, η
// οποία περιμένει να της υποδειχθεί τι πρέπει να κάνει.
func waitForTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("child : recv'd signal :", d)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"
	fmt.Println("parent : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------")
}
