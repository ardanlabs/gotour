package tour

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/blevesearch/bleve/v2"
)

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

	// -------------------------------------------------------------------------

	hitsIDs := make([]string, len(searchResults.Hits))
	for i, hit := range searchResults.Hits {
		hitsIDs[i] = hit.ID
	}
	sort.Strings(hitsIDs)

	// -------------------------------------------------------------------------

	result := make(map[string]lesson)

	for _, hit := range hitsIDs {
		lessonIDAndPage := strings.Split(hit, ".")

		// If the lessonID exists in the result, that means we already added
		// the pages. We can move to the next lesson ID.
		if _, exists := result[lessonIDAndPage[0]]; exists {
			continue
		}

		// Search for the lesson in the lessons.
		lsn, _ := lessons[lessonIDAndPage[0]]

		// -----------------------------------------------------------------------

		var pages []page

		// Iterate through the hits to find each page of a lesson and adds it
		// to the lesson in the right order.
		for _, h := range hitsIDs {
			hLessonIDAndPage := strings.Split(h, ".")

			// Check if the lessonID of the pages loop is the lessonID of the
			// lessons (parent) loop. If not, continue the loop.
			if lessonIDAndPage[0] != hLessonIDAndPage[0] {
				continue
			}

			pageNumber, err := strconv.Atoi(hLessonIDAndPage[1])
			if err != nil {
				log.Println(err)
			}

			pages = append(pages, lsn.Pages[pageNumber])
		}

		// -----------------------------------------------------------------------

		l := lesson{
			Title:       lsn.Title,
			Description: lsn.Description,
			Pages:       pages,
		}

		result[lessonIDAndPage[0]] = l
	}

	// -------------------------------------------------------------------------

	w.Header().Set("Content-Type", "application/json")

	if err := writeLessons(result, w); err != nil {
		err := fmt.Errorf("writing result lessons: %w", err)
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
