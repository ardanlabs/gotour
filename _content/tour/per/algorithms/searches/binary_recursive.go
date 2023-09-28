//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a binary search using a
// recursive approach.
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

	idx, err := binarySearchRecursive(numbers, numbers[find], 0, len(numbers))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Found  : Idx", idx)
}

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
