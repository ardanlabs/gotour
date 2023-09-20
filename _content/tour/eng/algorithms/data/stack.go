//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// This sample program shows you how to write a basic stack.
package main

import (
	"errors"
	"fmt"
)

func main() {
	const items = 5

	var s Stack

	for i := 0; i < items; i++ {
		name := fmt.Sprintf("Name%d", i)
		s.Push(Data{Name: name})

		fmt.Println("push:", name)
	}

	fmt.Println("------------------")

	f := func(d Data) error {
		fmt.Println("pop:", d.Name)
		return nil
	}

	s.Operate(f)
}

// Data represents what is being stored on the stack.
type Data struct {
	Name string
}

// Stack represents a stack of data.
type Stack struct {
	data []Data
}

// Make allows the creation of a stack with an initial
// capacity for efficiency. Otherwise a stack can be
// used in its zero value state.
func Make(cap int) *Stack {
	return &Stack{
		data: make([]Data, 0, cap),
	}
}

// Count returns the number of items in the stack.
func (s *Stack) Count() int {
	return len(s.data)
}

// Push adds data into the top of the stack.
func (s *Stack) Push(data Data) {
	s.data = append(s.data, data)
}

// Pop removes data from the top of the stack.
func (s *Stack) Pop() (Data, error) {
	if len(s.data) == 0 {
		return Data{}, errors.New("stack empty")
	}

	// Calculate the top level index.
	idx := len(s.data) - 1

	// Copy the data from that index position.
	data := s.data[idx]

	// Remove the top level index from the slice.
	s.data = s.data[:idx]

	return data, nil
}

// Peek provides the data stored on the stack based
// on the level from the bottom. A value of 0 would
// return the top piece of data.
func (s *Stack) Peek(level int) (Data, error) {
	if level < 0 || level > (len(s.data)-1) {
		return Data{}, errors.New("invalid level position")
	}
	idx := (len(s.data) - 1) - level
	return s.data[idx], nil
}

// Operate accepts a function that takes data and calls
// the specified function for every piece of data found.
// It traverses from the top down through the stack.
func (s *Stack) Operate(f func(data Data) error) error {
	for i := len(s.data) - 1; i > -1; i-- {
		if err := f(s.data[i]); err != nil {
			return err
		}
	}
	return nil
}
