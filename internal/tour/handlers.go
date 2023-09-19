package tour

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// rootHandler returns a handler for all the requests except the ones for lessons.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Get the cookies from the request
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

	if r.URL.Path == "/tour/eng/" {
		if err := renderUI(w); err != nil {
			log.Println(err)
		}
	}

	if r.URL.Path == "/tour/rus/" {
		if err := renderRusUI(w); err != nil {
			log.Println(err)
		}
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

// lessonRusHandler handler the HTTP requests for lessons.
func lessonRusHandler(w http.ResponseWriter, r *http.Request) {
	lesson := strings.TrimPrefix(r.URL.Path, "/tour/rus/lesson/")
	if err := writeRusLesson(lesson, w); err != nil {
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

	result, err := bleveSearch(index, lessons, qs)
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

// bleveRusHandler handler the HTTP requests for search.
func bleveRusHandler(w http.ResponseWriter, r *http.Request) {
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

	result, err := bleveSearch(rusIndex, rusLessons, qs)
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
