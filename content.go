// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package website exports the static content as an embed.FS.
package website

import (
	"embed"
	"io/fs"
)

//go:embed _content/favicon.ico
//go:embed _content/robots.txt
//go:embed _content/images/ardan-logo-dark.svg
//go:embed _content/images/favicon-gopher.svg
//go:embed _content/images/favicon-gopher-plain.png
//go:embed _content/js/playground.js
//go:embed _content/tour
var tourOnly embed.FS

// TourOnly returns the content needed only for the standalone tour.
func TourOnly() fs.FS {
	var fsys fs.FS = tourOnly
	s, err := fs.Sub(fsys, "_content")
	if err != nil {
		panic(err)
	}
	return s
}
