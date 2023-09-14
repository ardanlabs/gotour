//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create values from exported types with
// embedded unexported types.
package main

import (
	"fmt"

	"play.ground/users"
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

// -----------------------------------------------------------------------------
-- users/users.go --

// Package users provides support for user management.
package users

// User represents information about a user.
type user struct {
	Name string
	ID   int
}

// Manager represents information about a manager.
type Manager struct {
	Title string

	user
}

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
