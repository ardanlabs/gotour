// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tour

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
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
)

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

// =============================================================================

func Main() {
	httpListen = flag.String("http", "127.0.0.1:3999", "host:port to listen on")
	openBrowser = flag.Bool("openbrowser", true, "open browser automatically")
	webSocketOrigin = flag.String("origin", "", "host:port used for web socket origin")
	webSocketScheme = flag.String("scheme", "", "http or https, used for web socket origin scheme")

	flag.Parse()

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

	// -------------------------------------------------------------------------
	// Add Language Content

	engUIContent, engIndex := addLanguage("tour/eng/")
	defer engIndex.Close()

	grcUIContent, grcIndex := addLanguage("tour/grc/")
	defer grcIndex.Close()

	perUIContent, perIndex := addLanguage("tour/per/")
	defer perIndex.Close()

	porUIContent, porIndex := addLanguage("tour/por/")
	defer porIndex.Close()

	polUIContent, polIndex := addLanguage("tour/pol/")
	defer polIndex.Close()

	// -------------------------------------------------------------------------
	// Start Web Service

	r := root{
		engContent: engUIContent,
		grcContent: grcUIContent,
		perContent: perUIContent,
		porContent: porUIContent,
		polContent: polUIContent,
	}

	http.HandleFunc("/", r.rootHandler)
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

func addLanguage(route string) ([]byte, bleve.Index) {
	index, err := bleve.NewMemOnly(bleve.NewIndexMapping())
	if err != nil {
		log.Fatal(err)
	}

	routes := routes{
		index: index,
		route: route,
	}

	if err := initTour(http.DefaultServeMux, "SocketTransport", &routes); err != nil {
		log.Fatal(err)
	}

	return routes.uiContent, index
}

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

// =============================================================================

type logging struct {
	h http.Handler
}

func (l *logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	println(r.URL.Path)
	l.h.ServeHTTP(w, r)
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
