//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to use the predeclared type constraint
// "comparable". A type parameter with the comparable constraint accepts as
// a type argument any comparable type. It permits the use of == and != with
// values of that type parameter.
package main

import "fmt"

func index[T comparable](list []T, find T) int {
	for i, v := range list {
		if v == find {
			return i
		}
	}
	return -1
}

type person struct {
	name  string
	email string
}

func main() {
	durations := []int{5000, 10, 40}
	findDur := 10

	i := index(durations, findDur)
	fmt.Printf("Index: %d for %d\n", i, findDur)

	people := []person{
		{name: "bill", email: "bill@email.com"},
		{name: "jill", email: "jill@email.com"},
		{name: "tony", email: "tony@email.com"},
	}
	findPerson := person{name: "tony", email: "tony@email.com"}

	i = index(people, findPerson)
	fmt.Printf("Index: %d for %s\n", i, findPerson.name)
}
