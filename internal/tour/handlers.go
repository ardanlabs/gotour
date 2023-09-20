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
	dutContent []byte
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
			log.Println("redirect to language preference")
			http.Redirect(w, r, "/tour/"+langPref+"/", http.StatusFound)
			return
		}

		// Defaults to English version.
		log.Println("redirect to english")
		http.Redirect(w, r, "/tour/eng/", http.StatusFound)
		return
	}

	switch r.URL.Path {
	case "/tour/eng/":
		log.Println("render english tour")
		if err := renderUI(w, rot.engContent); err != nil {
			log.Println(err)
		}
	case "/tour/dut/":
		log.Println("render dutch tour")
		if err := renderUI(w, rot.dutContent); err != nil {
			log.Println(err)
		}
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
