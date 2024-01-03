//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to retrieve the minimum integer
// from a slice of integers.
package main

import "fmt"

func main() {
	tt := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{nil, 0},
		{[]int{10}, 10},
		{[]int{20, 30, 10, 50}, 10},
		{[]int{30, 50, 10}, 10},
	}

	for _, test := range tt {
		value, err := Min(test.input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Input: %d, Value: %d, Expected: %d, Match: %v\n",
			test.input,
			value,
			test.expected,
			value == test.expected,
		)
	}
}

// Min returns the minimum integer in the slice.
func Min(n []int) (int, error) {

	// First check there are numbers in the collection.
	if len(n) == 0 {
		return 0, fmt.Errorf("slice %#v has no elements", n)
	}

	// If the length of the slice is 1 then return the
	// integer at index 0.
	if len(n) == 1 {
		return n[0], nil
	}

	// Save the first value as current min and then loop over
	// the slice of integers looking for a smaller number.
	min := n[0]
	for _, num := range n[1:] {

		// If num is less than min. Assign min to num.
		if num < min {
			min = num
		}
	}

	return min, nil
}
