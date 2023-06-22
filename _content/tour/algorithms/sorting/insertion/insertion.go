//go:build OMIT
// +build OMIT

package main

func main() {

}
func insertionSort(randomList []int) []int {

	// Loop through the list until it is sorted.
	for leftIdx := 1; leftIdx < len(randomList); leftIdx++ {
		checkNum := randomList[leftIdx]
		rightIdx := leftIdx - 1

		// Look to check the number with the previous one. If the previous number is greater,
		// it will be swapped until it gets the correct position.
		for rightIdx >= 0 && randomList[rightIdx] > checkNum {
			randomList[rightIdx+1] = randomList[rightIdx]
			rightIdx--
		}

		randomList[rightIdx+1] = checkNum
	}

	return randomList
}
