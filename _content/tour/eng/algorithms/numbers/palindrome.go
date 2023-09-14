//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to determine if an integer is a
// palindrome or not.
package main

import "fmt"

func main() {
	tt := []int{-1, 1, 9, 10, 1001, 125}

	for _, input := range tt {
		success := IsPalindrome(input)

		switch success {
		case true:
			fmt.Printf("%d is a palindrome\n", input)

		case false:
			fmt.Printf("%d is NOT a palindrome\n", input)
		}
	}
}

// IsPalindrome checks if a integer is a Palindrome.
func IsPalindrome(input int) bool {

	// A negative integer is not a palindrome.
	if input < 0 {
		return false
	}

	// An integer that is only one digit in length is a palindrome.
	if input >= 0 && input < 10 {
		return true
	}

	// Reverse the digits in the integer.
	rev := Reverse(input)

	return input == rev
}

// Reverse takes the specified integer and reverses it.
func Reverse(num int) int {

	// Construct result to its zero value.
	var result int

	// Loop until num is zero.
	for num != 0 {

		// Perform a modulus operation to get the last digit from the value set in num.
		// https://www.geeksforgeeks.org/find-first-last-digits-number/
		// Ex. For num = 125, last = 5
		last := num % 10

		// Multiple the current result by 10 to shift the digits in
		// the current result to the left.
		// Ex. For result = 5, result = 50
		result = result * 10

		// Add the digit we took from the end of num to the result.
		// Ex. For result = 21 and last = 5, result = 215
		result += last

		// Remove the digit we just reversed from num.
		// Ex. For num = 125, num = 12
		num = num / 10
	}

	return result
}
