//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a bubble sort.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(10)
	fmt.Println("Before:", numbers)

	bubbleSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func bubbleSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n; i++ {
		if !sweep(numbers, i) {
			return
		}
	}
}

func sweep(numbers []int, currentPass int) bool {
	var idx int
	var swap bool

	idxNext := idx + 1
	n := len(numbers)

	for idxNext < (n - currentPass) {
		a := numbers[idx]
		b := numbers[idxNext]

		if a > b {
			numbers[idx] = b
			numbers[idxNext] = a
			swap = true
		}

		idx++
		idxNext = idx + 1
	}

	return swap
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
