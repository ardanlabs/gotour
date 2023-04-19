//go:build ignore || OMIT
// +build ignore OMIT

package main

import "github.com/ardanlabs/gotour/external/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
