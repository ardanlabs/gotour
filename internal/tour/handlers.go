package tour

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

type root struct {
	engContent []byte
	grcContent []byte
	perContent []byte
	porContent []byte
	polContent []byte
}

// rootHandler returns a handler for all the requests except the ones for lessons.
func (rot *root) rootHandler(w http.ResponseWriter, r *http.Request) {

	// Get the cookies from the request.
	cookies := r.Cookies()
	var langPref string

	// Iterate through the cookies to find the language preference cookie.
	for _, cookie := range cookies {
		if cookie.Name == "language-preference" {
			langPref = cookie.Value
			break
		}
	}

	if r.URL.Path == "/" {
		if langPref != "" {
			http.Redirect(w, r, "/tour/"+langPref+"/", http.StatusFound)
			return
		}

		// Defaults to English version.
		http.Redirect(w, r, "/tour/eng/", http.StatusFound)
		return
	}

	switch r.URL.Path {
	case "/tour/eng/":
		log.Println("render english tour")
		if err := renderUI(w, rot.engContent); err != nil {
			log.Println(err)
		}
	case "/tour/grc/":
		log.Println("render greek tour")
		if err := renderUI(w, rot.grcContent); err != nil {
			log.Println(err)
		}
	case "/tour/per/":
		log.Println("render persian tour")
		if err := renderUI(w, rot.perContent); err != nil {
			log.Println(err)
		}
	case "/tour/por/":
		log.Println("render portuguese tour")
		if err := renderUI(w, rot.porContent); err != nil {
			log.Println(err)
		}
	case "/tour/pol/":
		log.Println("render polish tour")
		if err := renderUI(w, rot.polContent); err != nil {
			log.Println(err)
		}
	default:
		http.Redirect(w, r, "/tour/eng/", http.StatusFound)
	}
}

// renderUI writes the tour UI to the provided Writer.
func renderUI(w io.Writer, content []byte) error {
	if content == nil {
		panic("renderUI called before successful initTour")
	}

	_, err := w.Write(content)
	return err
}

// =============================================================================

type routes struct {
	uiContent []byte
	index     bleve.Index
	lessons   map[string]lesson
	route     string
}

func (rou *routes) SetUIContent(content []byte) {
	rou.uiContent = content
}

func (rou *routes) SetLessons(lessons map[string]lesson) {
	rou.lessons = lessons
}

func (rou *routes) Route() string {
	return rou.route
}

func (rou *routes) Index() bleve.Index {
	return rou.index
}

func (rou *routes) Lessons() map[string]lesson {
	return rou.lessons
}

func (rou *routes) RootHandler(w http.ResponseWriter, r *http.Request) {
	if err := renderUI(w, rou.uiContent); err != nil {
		log.Println(err)
	}
}

func (rou *routes) LessonHandler(w http.ResponseWriter, r *http.Request) {
	lesson := strings.TrimPrefix(r.URL.Path, "/"+rou.route+"lesson/")
	if err := writeLesson(lesson, w, rou.lessons); err != nil {
		if err == ErrLessonNotFound {
			http.NotFound(w, r)
		} else {
			log.Println(err)
		}
	}
}

func (rou *routes) BleveHandler(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query().Get("search")
	qs = strings.Trim(qs, " ")

	// -------------------------------------------------------------------------

	if qs == "" {
		if _, err := fmt.Fprint(w, "{"); err != nil {
			log.Println(err)
		}

		if _, err := fmt.Fprint(w, "}"); err != nil {
			log.Println(err)
		}
		return
	}

	// -------------------------------------------------------------------------

	result, err := bleveSearch(rou.index, rou.lessons, qs)
	if err != nil {
		log.Println(err)
	}

	// -------------------------------------------------------------------------

	w.Header().Set("Content-Type", "application/json")

	if err := writeLessons(result, w); err != nil {
		err := fmt.Errorf("writing result lessons: %w", err)
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
