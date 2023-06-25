//go:build OMIT
// +build OMIT

package main

import (
	"fmt"
	"math/rand"
)

func main() {

}

// binarySearchRecursive takes the list of the sorted numbers and check it
// with `recursive` process to find the value and return the index of array or
// return the error if the value not found.
func binarySearchRecursive(sortedList []int, target int, leftIdx int, rightIdx int) (int, error) {

	// Calculate the middle index of the list.
	midIdx := (leftIdx + rightIdx) / 2

	// Check until leftIdx is smaller or equal with rightIdx.
	if leftIdx <= rightIdx {

		switch {

		// Check if we found the target.
		case sortedList[midIdx] == target:
			return midIdx, nil

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case sortedList[midIdx] > target:
			return binarySearchRecursive(sortedList, target, leftIdx, midIdx-1)

		// If the value is less than the target, cut the list
		// by moving the leftIdx into the list.
		case sortedList[midIdx] < target:
			return binarySearchRecursive(sortedList, target, midIdx+1, rightIdx)
		}
	}

	return -1, fmt.Errorf("target not found")
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
