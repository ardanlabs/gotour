//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a heap sort.
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numbers := generateList(10)
	fmt.Println("Before:", numbers)

	heapSort(numbers)
	fmt.Println("Sequential:", numbers)
}

func heapSort(numbers []int) []int {

	// Split the list in half and work the front half of the list, moving
	// the largest value we find to the front of the list and then the
	// second largest.

	for index := (len(numbers) / 2) - 1; index >= 0; index-- {
		numbers = moveLargest(numbers, len(numbers), index)
	}

	// Take the list and start moving numbers out and into a new sorted
	// list. Take the number in the first position and remove it to the
	// new list which will contain the final sort. Then move the largest
	// number we find once again to the front of the list.

	size := len(numbers)
	for index := size - 1; index >= 1; index-- {
		numbers[0], numbers[index] = numbers[index], numbers[0]
		size--
		numbers = moveLargest(numbers, size, 0)
	}

	return numbers
}

// moveLargest starts at the index positions specified in the list and attempts
// to move the largest number it can find to that position in the list.
func moveLargest(numbers []int, size int, index int) []int {

	// Calculate the index deviation so numbers in the list can be
	// compared and swapped if needed.
	// index 0: cmpIdx1: 1 cmpIdx2:  2   index 5: cmpIdx1: 11 cmpIdx2: 12
	// index 1: cmpIdx1: 3 cmpIdx2:  4   index 6: cmpIdx1: 13 cmpIdx2: 14
	// index 2: cmpIdx1: 5 cmpIdx2:  6   index 7: cmpIdx1: 15 cmpIdx2: 16
	// index 3: cmpIdx1: 7 cmpIdx2:  8   index 8: cmpIdx1: 17 cmpIdx2: 19
	// index 4: cmpIdx1: 9 cmpIdx2: 10   index 9: cmpIdx1: 19 cmpIdx2: 20
	cmpIdx1, cmpIdx2 := 2*index+1, 2*index+2

	// Save the specified index as the index with the current largest value.
	largestValueIdx := index

	// Check if the value at the first deviation index is greater than
	// the value at the current largest index. If so, save that
	// index position.
	if cmpIdx1 < size && numbers[cmpIdx1] > numbers[largestValueIdx] {
		largestValueIdx = cmpIdx1
	}

	// Check the second deviation index is within bounds and is greater
	// than the value at the current largest index. If so, save that
	// index position.
	if cmpIdx2 < size && numbers[cmpIdx2] > numbers[largestValueIdx] {
		largestValueIdx = cmpIdx2
	}

	// If we found a larger value than the value at the specified index, swap
	// those numbers and then recurse to find more numbers to swap from that
	// point in the list.
	if largestValueIdx != index {
		numbers[index], numbers[largestValueIdx] = numbers[largestValueIdx], numbers[index]
		numbers = moveLargest(numbers, size, largestValueIdx)
	}

	return numbers
}

func generateList(totalNumbers int) []int {
	numbers := make([]int, totalNumbers)

	for i := 0; i < totalNumbers; i++ {
		numbers[i] = rand.Intn(totalNumbers * 20)
	}

	return numbers
}
