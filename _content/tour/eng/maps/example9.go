//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how nil maps can be defined.
package main

import "fmt"

type user struct {
	name        string
	contactInfo map[string]string
}

func main() {
	u := user{}

	fmt.Printf("name value: 		`%s`\n", u.name)
	fmt.Printf("contactInfo value:	%v\n", u.contactInfo)
	fmt.Printf("is contactInfo nill:	%v\n", u.contactInfo == nil)
}
