// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to declare and use variadic functions.
package main

import "fmt"

// user is a struct type that declares user information.
type user struct {
	id   int
	name string
}

func main() {

	// Declare and initialize a value of type user.
	u1 := user{
		id:   1432,
		name: "Betty",
	}

	// Declare and initialize a value of type user.
	u2 := user{
		id:   4367,
		name: "Janet",
	}

	// Display both user values.
	display(u1, u2)

	// Create a slice of user values.
	u3 := []user{
		{24, "Bill"},
		{32, "Joan"},
	}

	// Display all the user values from the slice.
	display(u3...)

	change(u3...)
	fmt.Println("**************************")
	for _, u := range u3 {
		fmt.Printf("%+v\n", u)
	}
}

// display can accept and display multiple values of user types.
func display(users ...user) {
	fmt.Println("**************************")
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}
}

// change shows how the backing array is shared.
func change(users ...user) {
	users[1] = user{99, "Same Backing Array"}
}
