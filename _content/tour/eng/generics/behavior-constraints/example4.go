//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement a generics solution which allows
// a slice of some type T (to be determined later) to be passed and stringified.
// This code more closely resembles the concrete implementations that we started
// with and is easier to read than the reflect implementation. However, an
// interface constraint of type fmt.Stringer is applied to allow the compiler
// to know the value of type T passed requires a String method.
package main

import (
	"fmt"
)

func stringify[T fmt.Stringer](slice []T) []string {
	ret := make([]string, 0, len(slice))

	for _, value := range slice {
		ret = append(ret, value.String())
	}

	return ret
}

// Defining two types that implement the fmt.Stringer interface. Each
// implementation creates a stringified version of the concrete type.

type user struct {
	name  string
	email string
}

func (u user) String() string {
	return fmt.Sprintf("{type: \"user\", name: %q, email: %q}", u.name, u.email)
}

type customer struct {
	name  string
	email string
}

func (u customer) String() string {
	return fmt.Sprintf("{type: \"customer\", name: %q, email: %q}", u.name, u.email)
}

// =============================================================================

func main() {
	users := []user{
		{name: "Bill", email: "bill@ardanlabs.com"},
		{name: "Ale", email: "ale@whatever.com"},
	}

	s1 := stringify(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringify(customers)

	fmt.Println("customers:", s2)
}
