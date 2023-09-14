package tour

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// rootHandler returns a handler for all the requests except the ones for lessons.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/tour/eng/", http.StatusFound)
		return
	}
	if err := renderUI(w); err != nil {
		log.Println(err)
	}
}

// lessonHandler handler the HTTP requests for lessons.
func lessonHandler(w http.ResponseWriter, r *http.Request) {
	lesson := strings.TrimPrefix(r.URL.Path, "/tour/eng/lesson/")
	if err := writeLesson(lesson, w); err != nil {
		if err == ErrLessonNotFound {
			http.NotFound(w, r)
		} else {
			log.Println(err)
		}
	}
}

// bleveHandler handler the HTTP requests for search.
func bleveHandler(w http.ResponseWriter, r *http.Request) {
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

	result, err := bleveSearch(index, qs)
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
