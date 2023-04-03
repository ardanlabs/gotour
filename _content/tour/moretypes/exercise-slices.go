//go:build ignore || OMIT
// +build ignore OMIT

package main

import "github.com/ardanlabs/gotour/external/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
