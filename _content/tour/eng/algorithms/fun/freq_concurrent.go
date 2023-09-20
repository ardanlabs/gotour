// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample concurrent program shows you how to implement a function
// that can find the frequency a given rune is used in a specified sentence.
package main

import (
	"fmt"
	"runtime"
)

func main() {
	sentence := `The quick brown fox jumps over the lazy dog Stay hungry.
	Stay foolish Keep going. Be all in Boldness be my friend Screw it,
	let's do it My life is my message Leave no stone unturned Dream big.
	Pray bigger`

	print(concurrent(sentence))
}

func concurrent(text string) map[rune]int {
	m := make(map[rune]int)    // Map with final result
	g := runtime.GOMAXPROCS(0) // Number of goroutines
	l := len(text)             // Number of bytes to process
	b := l / g                 // Number of buckets, one per goroutine

	// Receives the result of each bucket processed
	// by a goroutine.
	ch := make(chan map[rune]int, g)

	// Create g number of goroutines.

	for i := 0; i < g; i++ {
		str := i * b   // Starting idx position of bucket
		end := str + b // Ending idx position of bucket
		if i == g-1 {  // The last bucket gets ant remaining bytes
			end = end + (l - end)
		}

		go func() {
			m := make(map[rune]int)

			defer func() {
				ch <- m
			}()

			// This G processes its bucket sequentially.
			for _, r := range text[str:end] {
				m[r]++
			}
		}()
	}

	// Wait for the results of each bucket to come
	// in and process them into the final map.

	for i := 0; i < g; i++ {
		result := <-ch
		for rk, rv := range result {
			m[rk] = m[rk] + rv
		}
	}

	return m
}

func print(m map[rune]int) {
	var cols int

	for r := 65; r < 65+26; r++ {
		v := m[rune(r)]
		fmt.Printf("%q:%d, ", rune(r), v)

		v = m[rune(r+32)]
		fmt.Printf("%q:%d, ", rune(r+32), v)

		cols++
		if cols == 5 {
			fmt.Print("\n")
			cols = 0
		}
	}
}
