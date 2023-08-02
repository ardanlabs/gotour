// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tour

import (
	"encoding/json"
	"flag"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/ardanlabs/gotour/internal/socket"
	"github.com/ardanlabs/gotour/internal/webtest"
	"github.com/blevesearch/bleve/v2"
)

const (
	socketPath = "/socket"
)

var (
	httpListen      *string
	openBrowser     *bool
	webSocketOrigin *string
	webSocketScheme *string

	httpAddr string
	scheme   string
	origin   string

	index bleve.Index
)

func Main() {

	httpListen = flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	openBrowser = flag.Bool("openbrowser", true, "open browser automatically")
	webSocketOrigin = flag.String("origin", "", "host:port used for web socket origin")
	webSocketScheme = flag.String("scheme", "", "http or https, used for web socket origin scheme")

	flag.Parse()

	// -------------------------------------------------------------------------
	// Create a new index mapping.

	var err error
	index, err = bleve.NewMemOnly(bleve.NewIndexMapping())
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// -------------------------------------------------------------------------

	host, port, err := net.SplitHostPort(*httpListen)
	if err != nil {
		log.Fatal(err)
	}
	if host == "" {
		host = "localhost"
	}
	if host != "127.0.0.1" && host != "localhost" {
		log.Print(localhostWarning)
	}

	// Used for when deploying to Cloud Run. The ENV var is set by Cloud Run as to the port to listen on.
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}

	// If origin scheme is specified we use that instead of http.
	scheme = "http"

	if *webSocketScheme != "" {
		scheme = *webSocketScheme
	}

	// If origin is specified we use that instead of host:port we are listening on.
	origin = host + ":" + port

	if *webSocketOrigin != "" {
		origin = *webSocketOrigin
	}

	httpAddr = host + ":" + port

	if err := initTour(http.DefaultServeMux, "SocketTransport", index); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/_/fmt", fmtHandler)
	fs := http.FileServer(http.FS(contentTour))
	http.Handle("/favicon.ico", fs)
	http.Handle("/robots.txt", fs)
	http.Handle("/images/", fs)

	// Specifies the origin scheme and host for web socket. For deployment they are different than running locally.
	originURL := &url.URL{Scheme: scheme, Host: origin}
	http.Handle(socketPath, socket.NewHandler(originURL))

	h := webtest.HandlerWithCheck(http.DefaultServeMux, "/_readycheck",
		os.DirFS("."), "tour/testdata/*.txt")

	go func() {
		url := "http://" + host + ":" + port
		if waitServer(url) && *openBrowser && startBrowser(url) {
			log.Printf("A browser window should open. If not, please visit %s", url)
		} else {
			log.Printf("Please open your web browser and visit %s", url)
		}
	}()

	log.Fatal(http.ListenAndServe(httpAddr, &logging{h}))
}

type logging struct {
	h http.Handler
}

func (l *logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	l.h.ServeHTTP(w, r)
}

func must(fsys fs.FS, err error) fs.FS {
	if err != nil {
		panic(err)
	}
	return fsys
}

// rootHandler returns a handler for all tfhe requests except the ones for lessons.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/tour/", http.StatusFound)
		return
	}
	if err := renderUI(w); err != nil {
		log.Println(err)
	}
}

// lessonHandler handler the HTTP requests for lessons.
func lessonHandler(w http.ResponseWriter, r *http.Request) {
	lesson := strings.TrimPrefix(r.URL.Path, "/tour/lesson/")
	if err := writeLesson(lesson, w); err != nil {
		if err == lessonNotFound {
			http.NotFound(w, r)
		} else {
			log.Println(err)
		}
	}
}

// lessonHandler handler the HTTP requests for lessons.
func bleveHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query().Get("search")

	// Create a query based on the user's search input.
	query := bleve.NewMatchPhraseQuery(qs)

	// Create a search request with the query.
	search := bleve.NewSearchRequest(query)

	// Perform the search on the bleve index.
	searchResults, err := index.Search(search)
	if err != nil {
		log.Println(err)
	}

	docIDs := make([]string, len(searchResults.Hits))
	for i, hit := range searchResults.Hits {
		docIDs[i] = hit.ID
	}

	data, err := json.Marshal(docIDs)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(data)
	if err != nil {
		log.Println(err)
	}
}

const localhostWarning = `
WARNING!  WARNING!  WARNING!

The tour server appears to be listening on an address that is
not localhost and is configured to run code snippets locally.
Anyone with access to this address and port will have access
to this machine as the user running gotour.

If you don't understand this message, hit Control-C to terminate this process.

WARNING!  WARNING!  WARNING!
`

// waitServer waits some time for the http Server to start
// serving url. The return value reports whether it starts.
func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(100 * time.Millisecond)
		tries--
	}
	return false
}

// startBrowser tries to open the URL in a browser, and returns
// whether it succeed.
func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

// prepContent for the local tour simply returns the content as-is.
var prepContent = func(r io.Reader) io.Reader { return r }

// socketAddr returns the WebSocket handler address.
var socketAddr = func() string {

	if scheme == "https" {
		return "wss://" + origin + socketPath
	}

	return "ws://tour.ardanlabs.com:443" + socketPath
}

// analyticsHTML is optional analytics HTML to insert at the beginning of <head>.
var analyticsHTML template.HTML
