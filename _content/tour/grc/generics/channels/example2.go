//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος προκειμένου να παρουσιαστεί ο τρόπος εκτέλεσης μιας συνάρτησης
// εργασίας μέσω ενός διαθέσιμου συνόλου από goroutine και η επιστροφή ενός καναλιού
// επικοινωνίας τύπου Input (που θα προσδιοριστεί αργότερα) πίσω στον καλώντα.
// Όταν η είσοδος παραληφθεί από οποιαδήποτε goroutine, η συνάρτηση εργασίας εκτελείται
// και η τιμή Result παρουσιάζεται.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

type poolWorkFn[Input any, Result any] func(input Input) Result

func poolWork[Input any, Result any](size int, work poolWorkFn[Input, Result]) (chan Input, func()) {
	var wg sync.WaitGroup
	wg.Add(size)

	ch := make(chan Input)

	for i := 0; i < size; i++ {
		go func() {
			defer wg.Done()
			for input := range ch {
				result := work(input)
				fmt.Println("pollWork :", result)
			}
		}()
	}

	cancel := func() {
		close(ch)
		wg.Wait()
	}

	return ch, cancel
}

func main() {
	size := runtime.GOMAXPROCS(0)
	pwf := func(input int) string {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return fmt.Sprintf("%d : received", input)
	}

	ch, cancel := poolWork(size, pwf)
	defer cancel()

	for i := 0; i < 5; i++ {
		ch <- i
	}
}
