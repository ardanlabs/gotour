//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement a stringify function that is
// specific to each of the concrete types implemented above. In each case,
// the stringify function returns a slice of strings. These function use
// the String method against the individual user or customer value.
package main

import (
	"fmt"
)

func stringifyUsers(users []user) []string {
	ret := make([]string, 0, len(users))
	for _, user := range users {
		ret = append(ret, user.String())
	}
	return ret
}

func stringifyCustomers(customers []customer) []string {
	ret := make([]string, 0, len(customers))
	for _, customer := range customers {
		ret = append(ret, customer.String())
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

	s1 := stringifyUsers(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyCustomers(customers)

	fmt.Println("customers:", s2)
}
