//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic stack type.
package main

import (
	"errors"
	"fmt"
)

// Declare a generic type named stack that uses a struct with a single
// field named data declared as a slice of some type T.
type stack[T any] struct {
	data []T
}

// Declare a method named push that accepts a value of some type T and appends
// the value to the slice.
func (s *stack[T]) push(v T) {
	s.data = append(s.data, v)
}

// Declare a method named pop that returns the latest value of some type T
// that was appended to the slice and an error.
func (s *stack[T]) pop() (T, error) {
	var zero T

	if len(s.data) == 0 {
		return zero, errors.New("stack is empty")
	}

	v := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return v, nil
}

// =============================================================================

func main() {

	// Constructs a value of type stack that stores integers.
	var s stack[int]

	// Push the values of 10 and 20 to the stack.
	s.push(10)
	s.push(20)

	// Pop a value from the stack.
	v, err := s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the value that was popped.
	fmt.Println(v)

	// Pop another value from the stack.
	v, err = s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the value that was popped.
	fmt.Println(v)

	// Pop another value from the stack. This should
	// return an error.
	v, err = s.pop()
	if err != nil {
		fmt.Println(err)
		return
	}
}
