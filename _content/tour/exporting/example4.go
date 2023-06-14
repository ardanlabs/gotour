// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// THE EXAMPLE WON'T RUN SINCE WE NEED PACKAGING

// =============================================================================
// users/users.go

// Package users provides support for user management.
package users

// User represents information about a user.
type User struct {
	Name string
	ID   int

	password string
}

// =============================================================================

// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/topics/go/language/exporting/example4/users"
)

func main() {

	// Create a value of type User from the users package.
	u := users.User{
		Name: "Chole",
		ID:   10,

		password: "xxxx",
	}

	// ./example4.go:21: unknown users.User field 'password' in struct literal

	fmt.Printf("User: %#v\n", u)
}
