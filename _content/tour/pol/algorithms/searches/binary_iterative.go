//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a binary search using an
// iterative approach.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := []int{4, 42, 80, 83, 121, 137, 169, 182, 185, 180}
	find := rand.Intn(10)

	fmt.Println("Numbers:", numbers)
	fmt.Println("Find.  :", numbers[find])

	idx, err := binarySearchIterative(numbers, numbers[find])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found  : Idx", idx)
}

func binarySearchIterative(sortedList []int, target int) (int, error) {
	var leftIdx int
	rightIdx := len(sortedList) - 1

	// Loop until we find the target or searched the list.
	for leftIdx <= rightIdx {

		// Calculate the middle index of the list.
		mid := (leftIdx + rightIdx) / 2

		// Capture the value to check.
		value := sortedList[mid]

		switch {

		// Check if we found the target.
		case value == target:
			return mid, nil

		// If the value is greater than the target, cut the list
		// by moving the rightIdx into the list.
		case value > target:
			rightIdx = mid - 1

		// If the value is less than the target, cut the list
		// by moving the leftIdx into the list.
		case value < target:
			leftIdx = mid + 1
		}
	}

	return -1, fmt.Errorf("target not found")
}
