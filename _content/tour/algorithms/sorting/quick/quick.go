package main

func main() {

}

func quickSort(randomList []int, leftIdx, rightIdx int) []int {
	switch {
	case leftIdx > rightIdx:
		return randomList

	// Divides array into two partitions.
	case leftIdx < rightIdx:
		randomList, pivotIdx := partition(randomList, leftIdx, rightIdx)

		quickSort(randomList, leftIdx, pivotIdx-1)
		quickSort(randomList, pivotIdx+1, rightIdx)
	}

	return randomList
}

// partition it takes a portion of an array then sort it.
func partition(randomList []int, leftIdx, rightIdx int) ([]int, int) {
	pivot := randomList[rightIdx]

	for smallest := leftIdx; smallest < rightIdx; smallest++ {
		if randomList[smallest] < pivot {
			randomList[smallest], randomList[leftIdx] = randomList[leftIdx], randomList[smallest]
			leftIdx++
		}
	}

	randomList[leftIdx], randomList[rightIdx] = randomList[rightIdx], randomList[leftIdx]

	return randomList, leftIdx
}
