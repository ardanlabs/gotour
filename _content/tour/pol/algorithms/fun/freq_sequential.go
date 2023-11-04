//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to implement a function that can
// find the frequency a given rune is used in a specified sentence.
package main

import (
	"fmt"
)

func main() {
	sentence := `The quick brown fox jumps over the lazy dog Stay hungry.
	Stay foolish Keep going. Be all in Boldness be my friend Screw it,
	let's do it My life is my message Leave no stone unturned Dream big.
	Pray bigger`

	print(sequential(sentence))
}

func sequential(text string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range text {
		m[r]++
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
