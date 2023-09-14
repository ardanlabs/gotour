//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
package main

import (
	"fmt"

	"play.ground/users"
)

func main() {

	// Create a value of type User from the users package.
	u := users.User{
		Name: "Chole",
		ID:   10,

		password: "xxxx",
	}

	// ./example4.go:21: unknown field password in struct literal of type users.User

	fmt.Printf("User: %#v\n", u)
}

// -----------------------------------------------------------------------------
-- users/users.go --

// Package users provides support for user management.
package users

// User represents information about a user.
type User struct {
	Name string
	ID   int

	password string
}

// -----------------------------------------------------------------------------
-- go.mod --
  
module "play.ground"

go 1.21.0
