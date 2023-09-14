//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a quick sort.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(5)
	fmt.Println("Before:", numbers)

	QuickSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func QuickSort(numbers []int) []int {
	return quickSort(numbers, 0, len(numbers)-1)
}

func quickSort(numbers []int, leftIdx, rightIdx int) []int {
	switch {
	case leftIdx > rightIdx:
		return numbers

	// Divides array into two partitions.
	case leftIdx < rightIdx:
		numbers, pivotIdx := partition(numbers, leftIdx, rightIdx)

		quickSort(numbers, leftIdx, pivotIdx-1)
		quickSort(numbers, pivotIdx+1, rightIdx)
	}

	return numbers
}

// partition it takes a portion of an array then sort it.
func partition(numbers []int, leftIdx, rightIdx int) ([]int, int) {
	pivot := numbers[rightIdx]

	for smallest := leftIdx; smallest < rightIdx; smallest++ {
		if numbers[smallest] < pivot {
			numbers[smallest], numbers[leftIdx] = numbers[leftIdx], numbers[smallest]
			leftIdx++
		}
	}

	numbers[leftIdx], numbers[rightIdx] = numbers[rightIdx], numbers[leftIdx]

	return numbers, leftIdx
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
