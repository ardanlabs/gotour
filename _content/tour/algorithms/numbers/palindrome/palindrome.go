//go:build OMIT
// +build OMIT

package main

func main() {

}

// Is checks if a integer is a Palindrome.
func Is(input int) bool {

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
