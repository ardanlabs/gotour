//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to mix type and behavior constraints.
package main

import "fmt"

// Defining two concrete types that implement a match method.

type person struct {
	name  string
	email string
}

func (p person) match(v person) bool {
	return p.name == v.name
}

type food struct {
	name     string
	category string
}

func (f food) match(v food) bool {
	return f.name == v.name
}

// The matcher interface defines two constraints. First, it constrains the data
// to what type is acceptable. Second, it constrains the behavior of the data.
// The match method requires that a value of type T (to be determined later)
// will be the input of the method.

// Note: The type list inside the interface is not needed for match to work.
//       I'm trying to show how the type list and behavior can be combined.

type matcher[T any] interface {
	person | food
	match(v T) bool
}

// The match function declares that the value of type T must implement the
// matcher interface and is used for the slice and value arguments to the
// function.

func match[T matcher[T]](list []T, find T) int {
	for i, v := range list {
		if v.match(find) {
			return i
		}
	}
	return -1
}

// =============================================================================

func main() {
	people := []person{
		{name: "bill", email: "bill@email.com"},
		{name: "jill", email: "jill@email.com"},
		{name: "tony", email: "tony@email.com"},
	}
	findPerson := person{name: "tony"}

	i := match(people, findPerson)
	fmt.Printf("Match: Idx: %d for %s\n", i, findPerson.name)

	foods := []food{
		{name: "apple", category: "fruit"},
		{name: "carrot", category: "veg"},
		{name: "chicken", category: "meat"},
	}
	findFood := food{name: "apple"}

	i = match(foods, findFood)
	fmt.Printf("Match: Idx: %d for %s\n", i, findFood.name)
}
