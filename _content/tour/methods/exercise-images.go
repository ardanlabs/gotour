//go:build ignore || OMIT
// +build ignore OMIT

package main

import "github.com/ardanlabs/gotour/external/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
