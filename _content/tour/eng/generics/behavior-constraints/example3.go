//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to implement a reflection solution which allows
// a slice of any type to be provided and stringified. This is a generic
// function thanks to the reflect package. Notice the call to the String
// method via reflection.
package main

import (
	"fmt"
	"reflect"
)

func stringifyReflect(v interface{}) []string {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Slice {
		return nil
	}

	ret := make([]string, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		m := val.Index(i).MethodByName("String")
		if !m.IsValid() {
			return nil
		}

		data := m.Call(nil)
		ret = append(ret, data[0].String())
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

	s1 := stringifyReflect(users)

	fmt.Println("users:", s1)

	// -------------------------------------------------------------------------

	customers := []customer{
		{name: "Google", email: "you@google.com"},
		{name: "MSFT", email: "you@msft.com"},
	}

	s2 := stringifyReflect(customers)

	fmt.Println("customers:", s2)
}
