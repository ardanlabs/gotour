//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Implement a generic function named marshal that can marshal JSON but only
// accepts values that implement the json.Marshaler interface.
package main

import (
	"encoding/json"
	"fmt"
)

// Implement the generic function named marshal that can accept only values
// of type T that implement the json.Marshaler interface.
func marshal[T json.Marshaler](v T) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// =============================================================================

// Define a type names user with two fields, name and email.
type user struct {
	name  string
	email string
}

// Declare a method that implements the json.Marshaler interface. Have the
// method return a value of type user as JSON.
func (u user) MarshalJSON() ([]byte, error) {
	v := fmt.Sprintf("{\"name\": %q, \"email\": %q}", u.name, u.email)
	return []byte(v), nil
}

// =============================================================================

func main() {

	// Construct a value of type user.
	user := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Call the generic marshal function.
	s1, err := marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display the returned JSON.
	fmt.Println("user:", string(s1))
}
