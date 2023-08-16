// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tour

import (
	"log"
	"net/http"
	"testing"

	"github.com/ardanlabs/gotour/internal/webtest"
	"github.com/blevesearch/bleve/v2"
)

func TestWeb(t *testing.T) {
	var err error
	index, err = bleve.NewMemOnly(bleve.NewIndexMapping())
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	if err := initTour(http.DefaultServeMux, "SocketTransport", index); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", rootHandler)
	webtest.TestHandler(t, "testdata/*.txt", http.DefaultServeMux)
}
