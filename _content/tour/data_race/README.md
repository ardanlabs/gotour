## Data Races

A data race is when two or more goroutines attempt to read and write to the same resource at the same time. Race conditions can create bugs that appear totally random or can never surface as they corrupt data. Atomic functions and mutexes are a way to synchronize the access of shared resources between goroutines.

## Notes

* Goroutines need to be coordinated and synchronized.
* When two or more goroutines attempt to access the same resource, we have a data race.
* Atomic functions and mutexes can provide the support we need.

## Cache Coherency and False Sharing
This content is provided by Scott Meyers from his talk in 2014 at Dive:

[CPU Caches and Why You Care (30:09-38:30)](https://youtu.be/WDIkqP4JbkE?t=1809)  
[Code Example](../../testing/benchmarks/falseshare/README.md)

![figure1](figure1.png)

## Cache Coherency and False Sharing Notes

* Thread memory access matters.
* If your algorithm is not scaling look for false sharing problems.

## Links

[Eliminate False Sharing](http://www.drdobbs.com/parallel/eliminate-false-sharing/217500206) - Herb Sutter    
[The Go Memory Model](https://golang.org/ref/mem)    
[Introducing the Go Race Detector](http://blog.golang.org/race-detector) - Dmitry Vyukov and Andrew Gerrand    
[Detecting Race Conditions With Go](https://www.ardanlabs.com/blog/2013/09/detecting-race-conditions-with-go.html) - William Kennedy    
[Data Race Detector](https://golang.org/doc/articles/race_detector.html)    

## Diagram

### View of Data Race in Example1.

![Ardan Labs](data_race.png)

## Code Review

[Data Race](example1/example1.go) ([Go Playground](https://play.golang.org/p/czqXM5wOspX))    
[Atomic Increments](example2/example2.go) ([Go Playground](https://play.golang.org/p/5ZtLaX7zxt7))    
[Mutex](example3/example3.go) ([Go Playground](https://play.golang.org/p/ZKE2v9H4oS-))    
[Read/Write Mutex](example4/example4.go) ([Go Playground](https://play.golang.org/p/-iXzElPBnDM))    
[Map Data Race](example5/example5.go) ([Go Playground](https://play.golang.org/p/ktWRjcJWNjw))    

## Advanced Code Review

[Interface Based Race Condition](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/fwRTeBQrZVW))

## Exercises

### Exercise 1
Given the following program, use the race detector to find and correct the data race.

	// https://play.golang.org/p/F5DCJTZ6Lm

	// Fix the race condition in this program.
	package main

	import (
		"fmt"
		"math/rand"
		"sync"
		"time"
	)

	// numbers maintains a set of random numbers.
	var numbers []int

	// init is called prior to main.
	func init() {
		rand.Seed(time.Now().UnixNano())
	}

	// main is the entry point for the application.
	func main() {
		// Number of goroutines to use.
		const grs = 3

		// wg is used to manage concurrency.
		var wg sync.WaitGroup
		wg.Add(grs)

		// Create three goroutines to generate random numbers.
		for i := 0; i < grs; i++ {
			go func() {
				random(10)
				wg.Done()
			}()
		}

		// Wait for all the goroutines to finish.
		wg.Wait()

		// Display the set of random numbers.
		for i, number := range numbers {
			fmt.Println(i, number)
		}
	}

	// random generates random numbers and stores them into a slice.
	func random(amount int) {
		// Generate as many random numbers as specified.
		for i := 0; i < amount; i++ {
			n := rand.Intn(100)
			numbers = append(numbers, n)
		}
	}

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/Mzt11_xe_ou)) | 
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/KAakUVF_1k-))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
