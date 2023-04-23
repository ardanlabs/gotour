package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	var filePath string

	flag.StringVar(
		&filePath,
		"file_path",
		"./_content/tour/pointers.article",
		"help message for flagname",
	)

	flag.Parse()

	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	r := regexp.MustCompile(`\[([\s\S]+?)\]\(([\s\S]+?)\)`)
	rTitle := regexp.MustCompile(`\[([\s\S]+?)\]`)
	rURL := regexp.MustCompile(`\(([\s\S]+?)\)`)

	matches := r.FindAllString(string(data), -1)
	document := string(data)

	for _, m := range matches {

		title := rTitle.FindStringSubmatch(m)[1]
		url := rURL.FindStringSubmatch(m)[1]

		document = strings.ReplaceAll(
			document,
			m,
			fmt.Sprintf(
				"[[%s][%s]]",
				url,
				title,
			),
		)
	}

	err = os.WriteFile(
		filePath,
		[]byte(document),
		os.ModePerm,
	)

	if err != nil {
		panic(err)
	}

}
