package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// Path where you want to store the index.
	indexFolderPath := "x_indexes/" + time.Now().Format("150405")

	// TODO: reuse the latest created index mapping.
	// Create a new index mapping.
	indexMapping := bleve.NewIndexMapping()

	index, err := bleve.New(indexFolderPath, indexMapping)
	if err != nil {
		panic(err)
	}
	defer index.Close()

	// -------------------------------------------------------------------------

	// Path to the folder containing your the files.
	contentFilesPath := "_content_test"

	// Index the content files in the specified folder.
	err = indexFiles(index, contentFilesPath)
	if err != nil {
		panic(err)
	}

	// -------------------------------------------------------------------------

	// User input for the search query.
	var searchQuery string
	fmt.Print("Enter search query: ")
	if _, err := fmt.Scanln(&searchQuery); err != nil {
		panic(err)
	}

	// -------------------------------------------------------------------------

	// Create a query based on the user's search input.
	query := bleve.NewQueryStringQuery(searchQuery)

	// Create a search request with the query.
	search := bleve.NewSearchRequest(query)

	// Perform the search on the bleve index.
	searchResults, err := index.Search(search)
	if err != nil {
		panic(err)
	}

	// Print the search results.
	for _, hit := range searchResults.Hits {
		// TODO: Investigate the other fields on the results.
		fmt.Printf("Document ID: %s\n", hit.ID)
	}

	fmt.Printf("Total hits: %d\n", searchResults.Total)
}

func indexFiles(index bleve.Index, folderPath string) error {

	// Walk through the specified folder and its subdirectories.
	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current item is not a directory and has the ".article"
		// extension.
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".article") {

			// Read the content of the file.
			htmlContent, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Create a unique document ID using the file name without the ".article"
			// extension.
			docID := strings.TrimSuffix(info.Name(), ".article")

			// TODO: I need to understand better this concept of documents and if
			//  this is better.
			// Create a new Bleve document and add a field with the content of the file
			// doc := document.NewDocument(docID)
			// doc.AddField(document.NewTextField("content", []uint64{}, htmlContent))

			// Index the document in the Bleve index.
			err = index.Index(docID, string(htmlContent))
			if err != nil {
				return err
			}
		}

		return nil
	}

	return filepath.Walk(folderPath, walkFunc)
}
