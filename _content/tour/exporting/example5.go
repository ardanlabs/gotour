// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// THESE EXAMPLES WON'T RUN SINCE WE NEED PACKAGING

// =============================================================================
// users/users.go

// Package users provides support for user management.
package users

// user represents information about a user.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
type Manager struct {
	Title string

	user
}

// =============================================================================

// Sample program to show how to create values from exported types with
// embedded unexported types.
package main

import (
	"fmt"

	"github.com/ardanlabs/gotraining/topics/go/language/exporting/example5/users"
)

func main() {

	// Create a value of type Manager from the users package.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// Set the exported fields from the unexported user inner type.
	u.Name = "Chole"
	u.ID = 10

	fmt.Printf("User: %#v\n", u)
}
