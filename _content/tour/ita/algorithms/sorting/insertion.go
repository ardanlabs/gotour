//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a insertion sort.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(10)
	fmt.Println("Before:", numbers)

	insertionSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func insertionSort(numbers []int) {
	var n = len(numbers)

	// Walk through the numbers from left to right. Through
	// each outer loop iteration we move values from right
	// to left inside the array when they are larger than
	// the value that preceed it.

	for i := 1; i < n; i++ {
		j := i

		// For the given starting i index position, look
		// for smaller values to move left down the numbers list.

		for j > 0 {

			// Is the value on the left larger than the
			// right. If true, swap the two values.

			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}

			// Walk through the item from right to left.

			j--
		}
	}
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
