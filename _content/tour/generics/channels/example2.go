//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to execute a work function via a pool of goroutines
// and return a channel of type Input (to be determined later) back to the caller.
// Once input is received by any given goroutine, the work function is executed
// and the Result value is displayed.
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
