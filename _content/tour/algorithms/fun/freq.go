//go:build OMIT
// +build OMIT

// Package freq provides support for find the frequency in which a rune
// is found in a collection of text documents.
package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {
	sentence := buildText(`The quick brown fox jumps over the lazy dog Stay hungry. Stay
	foolish Keep going. Be all in Boldness be my friend Screw it, let’s do it My
	life is my message Leave no stone unturned Dream big. Pray bigger`)

	// Sequential
	seq_result := Sequential(sentence)
	log.Println("Sequantial Output:", seq_result)

	// =============================================================================
	// Concurrent Bounded

	con_bounded := ConcurrentBounded(sentence)
	log.Println("Concurrent Bounded Output:", con_bounded)

	// =============================================================================
	// Concurrent Bounded Channel
	con_bounded_channel := ConcurrentBoundedChannel(sentence)
	log.Println("Concurrent Bounded Output:", con_bounded_channel)

	// =============================================================================
	//Concurrent Unlimited
	con_bounded_unlmtd := ConcurrentUnlimited(sentence)
	log.Println("Concurrent Bounded Output:", con_bounded_unlmtd)
}

func buildText(sentence string) []string {
	const n = 100

	var out = map[rune]int{
		'T': 1, 'q': 1, 'p': 2, '’': 1, 'i': 11, 'b': 4, 'w': 2, 'j': 1, 'B': 2,
		'L': 1, 'e': 20, 'v': 2, 'l': 7, ',': 1, 'h': 4, 'u': 5, 'f': 4, 's': 9,
		'g': 8, 'D': 1, 'P': 1, ' ': 37, 'z': 1, 'd': 5, '.': 3, 'c': 2, 'r': 9,
		'o': 11, 'm': 5, '\n': 2, 'x': 1, 'y': 8, 'S': 3, 'K': 1, 'k': 1, 'n': 10,
		't': 8, 'a': 8, 'M': 1,
	}

	var s []string
	for i := 0; i < n; i++ {
		s = append(s, sentence)
	}
	for k, v := range out {
		out[k] = v * n
	}
	return s
}

// Sequential uses a sequential algorithm.
func Sequential(text []string) map[rune]int {
	m := make(map[rune]int)
	for _, words := range text {
		for _, r := range words {
			m[r]++
		}
	}
	return m
}

// ConcurrentUnlimited uses a concurrent algorithm based on an
// unlimited fan out pattern.
func ConcurrentUnlimited(text []string) map[rune]int {
	ch := make(chan map[rune]int, len(text))
	for _, words := range text {
		go func(words string) {
			lm := make(map[rune]int)
			for _, r := range words {
				lm[r]++
			}
			ch <- lm
		}(words)
	}

	all := make(map[rune]int)
	for range text {
		lm := <-ch
		for r, c := range lm {
			all[r] += c
		}
	}

	return all
}

// ConcurrentBounded uses a concurrent algorithm based on a bounded
// fan out and no channels.
func ConcurrentBounded(text []string) map[rune]int {
	m := make(map[rune]int)

	goroutines := runtime.GOMAXPROCS(0)
	totalNumbers := len(text)
	lastGoroutine := goroutines - 1
	stride := totalNumbers / goroutines

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {
			lm := make(map[rune]int)
			defer func() {
				mu.Lock()
				defer mu.Unlock()
				for k, v := range lm {
					m[k] = m[k] + v
				}
				wg.Done()
			}()

			start := g * stride
			end := start + stride
			if g == lastGoroutine {
				end = totalNumbers
			}

			for _, words := range text[start:end] {
				for _, r := range words {
					lm[r]++
				}
			}
		}(g)
	}

	wg.Wait()
	return m
}

// ConcurrentBoundedChannel uses a concurrent algorithm based on a bounded
// fan out using a channel.
func ConcurrentBoundedChannel(text []string) map[rune]int {
	m := make(map[rune]int)

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	var mu sync.Mutex
	ch := make(chan string, len(text))

	for i := 0; i < g; i++ {
		go func() {
			lm := make(map[rune]int)
			defer func() {
				mu.Lock()
				defer mu.Unlock()
				for k, v := range lm {
					m[k] = m[k] + v
				}
				wg.Done()
			}()

			for words := range ch {
				for _, r := range words {
					lm[r]++
				}
			}
		}()
	}

	for _, words := range text {
		ch <- words
	}
	close(ch)

	wg.Wait()
	return m
}
