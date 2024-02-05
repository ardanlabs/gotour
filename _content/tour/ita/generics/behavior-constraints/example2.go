//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement an empty interface solution which
// uses type assertions for the different concrete slices to be supported.
// We've basically moved the functions from above into case statements.
// This function uses the String method against the value.
package main

import (
	"fmt"
)

func stringifyAssert(v interface{}) []string {
	switch list := v.(type) {
	case []user:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret

	case []customer:
		ret := make([]string, 0, len(list))
		for _, value := range list {
			ret = append(ret, value.String())
		}
		return ret
	}

	return nil
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

	s1 := stringifyAssert(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyAssert(customers)

	fmt.Println("customers:", s2)
}
