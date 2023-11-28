/*
	When adding a new language to the system, please follow the instructions below carefully:

	## GO Files
	- internal/tour/handlers.go:13 : Add a field for the new language ui content.
	- internal/tour/handlers.go:46 : Incorporate the new language into the switch/case statement to handle its routing properly.
	- internal/tour/local.go:101   : Register the new language by adding a new call to addLanguage function.
	- internal/tour/local.go:121   : Add the ui content to the root struct.

	## _content Folder
	1. Create a new language folder:
     - Make a copy of the "eng" folder under "_content/tour". Use three-letter code for the new language as the folder name (e.g., "esp" for Spanish).

	2. Adjust Static Content:
	   - _content/tour/[new_lang_code]/static: Replace all entries containing `tour/eng` with `tour/[new_lang_code]`.

	3. Update JavaScript Functions:
	   - _content/tour/[new_lang_code]/static/js/page.js (function: replaceLanguageInUrl): Add the new language code to the regex. Repeat this step in each language folder.
	   - _content/tour/[new_lang_code]/static/js/values.js: Replace the table of contents with the translated content for the new language.

	4. Modify Template:
	   - _content/tour/[new_lang_code]/template/*.tmpl        : Replace all entries containing `tour/eng` with `tour/[new_lang_code]`.
	   - _content/tour/[new_lang_code]/template/index.tmpl:42 : Translate the title to the new language.
	   - _content/tour/[new_lang_code]/template/index.tmpl:71 : Add the new language option to the language selection dropdown. Repeat this step in each language folder's template file.

	## Integrating New Language
	1. Update replace statement in all languages
	   - _content/tour/[new_lang_code]/static/js/page.js:116 : return url.replace(/(\/tour\/)(eng|rus|per|grc|por|[new_lang_code])(\/)/, `$1${newLanguage}$3`);

	2. Change Patch Number in all JS and Template files for all languages
	   - _content/tour/[new_lang_code]/welcome.article
	   - _content/tour/[new_lang_code]/static/partials/list.html
	   - _content/tour/[new_lang_code]/template/index.tmpl

	3. Update Language switch in all languages
	   - _content/tour/[new_lang_code]/template/index.tmpl:71 : <option value="[new_lang_code]">[new_lang]</option>

	## Testing
	- After completing the above steps, thoroughly test the new language integration to ensure that all aspects function correctly and the content appears as expected.
*/

package tour

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	website "github.com/ardanlabs/gotour"
	"github.com/blevesearch/bleve/v2"
	"golang.org/x/tools/present"
)

var ErrLessonNotFound = fmt.Errorf("lesson not found")

var contentTour = website.TourOnly()

type handler interface {
	SetUIContent(content []byte)
	SetLessons(lessons map[string]lesson)
	Route() string
	Index() bleve.Index
	Lessons() map[string]lesson
	RootHandler(w http.ResponseWriter, r *http.Request)
	LessonHandler(w http.ResponseWriter, r *http.Request)
	BleveHandler(w http.ResponseWriter, r *http.Request)
}

// file defines the JSON form of a code file in a page.
type file struct {
	Name    string
	Content string
	Hash    string
}

// page defines the JSON form of a tour lesson page.
type page struct {
	Title   string
	Content string
	Files   []file
}

// lesson defines the JSON form of a tour lesson.
type lesson struct {
	Title       string
	Description string
	Pages       []page
}

// initTour loads tour.article, relevant HTML templates from root, and
// initialize the bleve index.
func initTour(mux *http.ServeMux, transport string, h handler) error {

	// Make sure playground is enabled before rendering.
	present.PlayEnabled = true

	// -------------------------------------------------------------------------
	// Content

	// Set up templates.
	tmpl, err := present.Template().ParseFS(contentTour, h.Route()+"template/action.tmpl")
	if err != nil {
		return fmt.Errorf("parse %s templates: %v", h.Route(), err)
	}

	// Init lessons.
	location := h.Route()[:len(h.Route())-1] // Remove the leading /
	files, err := fs.ReadDir(contentTour, location)
	if err != nil {
		return err
	}
	if err := initLessons(tmpl, files, h); err != nil {
		return fmt.Errorf("init %s lessons: %v", h.Route(), err)
	}

	// Index lessons into the bleve index.
	if err := indexLessons(h.Index(), h.Lessons()); err != nil {
		return fmt.Errorf("indexing %s lessons: %v", h.Route(), err)
	}

	// -------------------------------------------------------------------------
	// Init UI

	ui, err := template.ParseFS(contentTour, h.Route()+"template/index.tmpl")
	if err != nil {
		return fmt.Errorf("parse %s index.tmpl: %v", h.Route(), err)
	}
	buf := new(bytes.Buffer)

	data := struct {
		AnalyticsHTML template.HTML
	}{analyticsHTML}

	if err := ui.Execute(buf, data); err != nil {
		return fmt.Errorf("render %s UI: %v", h.Route(), err)
	}
	h.SetUIContent(buf.Bytes())

	// -------------------------------------------------------------------------
	// Set language specific routes.

	mux.HandleFunc("/"+h.Route(), h.RootHandler)
	mux.HandleFunc("/"+h.Route()+"lesson/", h.LessonHandler)
	mux.HandleFunc("/"+h.Route()+"bleve/", h.BleveHandler)
	mux.Handle("/"+h.Route()+"static/", http.FileServer(http.FS(contentTour)))

	// -------------------------------------------------------------------------

	return initScript(mux, socketAddr(), transport, h.Route())
}

// initScript concatenates all the javascript files needed to render
// the tour UI and serves the result on /script.js.
func initScript(mux *http.ServeMux, socketAddr, transport string, route string) error {
	modTime := time.Now()
	b := new(bytes.Buffer)

	// Keep this list in dependency order
	files := []string{
		"../../js/playground.js",
		"static/lib/jquery.min.js",
		"static/lib/jquery-ui.min.js",
		"static/lib/angular.min.js",
		"static/lib/codemirror/lib/codemirror.js",
		"static/lib/codemirror/mode/go/go.js",
		"static/lib/angular-ui.min.js",
		"static/js/app.js",
		"static/js/controllers.js",
		"static/js/directives.js",
		"static/js/services.js",
		"static/js/values.js",
	}

	for _, file := range files {
		f, err := fs.ReadFile(contentTour, path.Clean(route+file))
		if err != nil {
			return err
		}
		b.Write(f)
	}

	f, err := fs.ReadFile(contentTour, route+"static/js/page.js")
	if err != nil {
		return err
	}
	s := string(f)
	s = strings.ReplaceAll(s, "{{.SocketAddr}}", socketAddr)
	s = strings.ReplaceAll(s, "{{.Transport}}", transport)
	b.WriteString(s)

	mux.HandleFunc("/"+route+"script.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/javascript")
		// Set expiration time in one week.
		w.Header().Set("Cache-control", "max-age=604800")
		http.ServeContent(w, r, "", modTime, bytes.NewReader(b.Bytes()))
	})

	return nil
}

// initEngLessons finds all the lessons in the content directory, renders them,
// using the given template and saves the content in the lessons map.
func initLessons(tmpl *template.Template, files []fs.DirEntry, h handler) error {
	lessons := make(map[string]lesson)

	for _, f := range files {
		if path.Ext(f.Name()) != ".article" {
			continue
		}

		lsn, err := parseLesson(h.Route(), f.Name(), tmpl)
		if err != nil {
			return fmt.Errorf("parsing %v: %v", f.Name(), err)
		}

		name := strings.TrimSuffix(f.Name(), ".article")

		w := new(bytes.Buffer)
		if err := json.NewEncoder(w).Encode(lsn); err != nil {
			return fmt.Errorf("encode lesson: %v", err)
		}

		lessons[name] = lsn
	}

	h.SetLessons(lessons)

	return nil
}

// parseLesson parses and returns a lesson content given its path
// relative to root ('/'-separated) and the template to render it.
func parseLesson(directory string, path string, tmpl *template.Template) (lesson, error) {
	f, err := contentTour.Open(directory + path)
	if err != nil {
		return lesson{}, err
	}
	defer f.Close()
	ctx := &present.Context{
		ReadFile: func(filename string) ([]byte, error) {
			return fs.ReadFile(contentTour, directory+filepath.ToSlash(filename))
		},
	}
	doc, err := ctx.Parse(prepContent(f), path, 0)
	if err != nil {
		return lesson{}, err
	}

	lsn := lesson{
		doc.Title,
		doc.Subtitle,
		make([]page, len(doc.Sections)),
	}

	for i, sec := range doc.Sections {
		p := &lsn.Pages[i]
		w := new(bytes.Buffer)
		if err := sec.Render(w, tmpl); err != nil {
			return lesson{}, fmt.Errorf("render section: %v", err)
		}
		p.Title = sec.Title
		p.Content = w.String()
		codes := findPlayCode(sec)
		p.Files = make([]file, len(codes))
		for i, c := range codes {
			f := &p.Files[i]
			f.Name = c.FileName
			f.Content = string(c.Raw)
			hash := sha1.Sum(c.Raw)
			f.Hash = base64.StdEncoding.EncodeToString(hash[:])
		}
	}

	return lsn, nil
}

// findPlayCode returns a slide with all the Code elements in the given
// Elem with Play set to true.
func findPlayCode(e present.Elem) []*present.Code {
	var r []*present.Code
	switch v := e.(type) {
	case present.Code:
		if v.Play {
			r = append(r, &v)
		}
	case present.Section:
		for _, s := range v.Elem {
			r = append(r, findPlayCode(s)...)
		}
	}
	return r
}

// writeLesson writes the tour content to the provided Writer.
func writeLesson(name string, w io.Writer, lessons map[string]lesson) error {
	if len(name) == 0 {
		return writeAllLessons(w, lessons)
	}

	l, ok := lessons[name]
	if !ok {
		return ErrLessonNotFound
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(l); err != nil {
		log.Printf("encode lesson: %v", err)
	}

	_, err := w.Write(b.Bytes())
	return err
}

func writeAllLessons(w io.Writer, lessons map[string]lesson) error {
	if _, err := fmt.Fprint(w, "{"); err != nil {
		return err
	}

	nLessons := len(lessons)
	for k, v := range lessons {
		b := new(bytes.Buffer)
		if err := json.NewEncoder(b).Encode(v); err != nil {
			log.Printf("encode lesson: %v", err)
		}

		if _, err := fmt.Fprintf(w, "%q:%s", k, b.Bytes()); err != nil {
			return err
		}

		nLessons--
		if nLessons != 0 {
			if _, err := fmt.Fprint(w, ","); err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprint(w, "}")
	return err
}

func writeLessons(l map[string]lesson, w io.Writer) error {
	if _, err := fmt.Fprint(w, "{"); err != nil {
		return err
	}

	nLessons := len(l)
	for k, v := range l {
		b := new(bytes.Buffer)
		if err := json.NewEncoder(b).Encode(v); err != nil {
			log.Printf("encode lesson: %v", err)
		}

		if _, err := fmt.Fprintf(w, "%q:%s", k, b.Bytes()); err != nil {
			return err
		}

		nLessons--
		if nLessons != 0 {
			if _, err := fmt.Fprint(w, ","); err != nil {
				return err
			}
		}
	}

	_, err := fmt.Fprint(w, "}")
	return err
}

// indexLessons initializes the provided bleve index with content from lessons.
// It iterates through each lesson's pages, excluding "Exercises" pages,
// and indexes the content using a formatted ID that combines the lesson
// name and page number. The content is structured as an ID-Content pair.
func indexLessons(index bleve.Index, lessons map[string]lesson) error {
	for lessonName, lsn := range lessons {
		if err := indexLesson(index, lessonName, lsn); err != nil {
			return fmt.Errorf("failed to index rus lesson %s: %w", lessonName, err)
		}
	}

	return nil
}

// indexLesson indexes the pages of a lesson into the provided bleve index.
func indexLesson(index bleve.Index, lessonName string, lsn lesson) error {
	for pageNum, page := range lsn.Pages {

		// Skip indexing "Exercise" pages.
		if strings.Contains(page.Title, "Exercise") {
			continue
		}

		contentID := fmt.Sprintf("%s.%d", lessonName, pageNum)

		data := struct {
			ID      string
			Content string
		}{
			ID:      contentID,
			Content: page.Content,
		}

		if err := index.Index(contentID, data); err != nil {
			return fmt.Errorf("failed to index content %s: %w", contentID, err)
		}
	}

	return nil
}

// bleveSearch performs a search on the provided bleve index using the given
// match phrase.
// It creates a query based on the match phrase, performs the search, and
// organizes the search results into a map of lessons and their relevant pages.
func bleveSearch(index bleve.Index, lessons map[string]lesson, matchPhrase string) (map[string]lesson, error) {
	// Create a query based on the user's search input.
	query := bleve.NewMatchPhraseQuery(matchPhrase)

	// Create a search request with the query.
	search := bleve.NewSearchRequest(query)

	// Perform the search on the bleve index.
	searchResults, err := index.Search(search)
	if err != nil {
		return nil, err
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
		lsn := lessons[lessonIDAndPage[0]]

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
				return nil, err
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

	return result, nil
}
