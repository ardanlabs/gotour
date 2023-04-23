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
	var modName string

	flag.StringVar(
		&filePath,
		"file_path",
		"./_content/tour/pointers.article",
		"Path to Markdown file",
	)

	flag.StringVar(
		&modName,
		"module_name",
		"pointers",
		"name of module",
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
	document = strings.ReplaceAll(
		document,
		"*",
		"-",
	)

	document = strings.ReplaceAll(
		document,
		"### ",
		"",
	)

	document = strings.ReplaceAll(
		document,
		"## ",
		"* ",
	)

	for _, m := range matches {

		title := rTitle.FindStringSubmatch(m)[1]
		url := rURL.FindStringSubmatch(m)[1]

		if strings.Contains(url, "https://play.golang.org") {
			document = strings.ReplaceAll(
				document,
				fmt.Sprintf("(%s)", m),
				"",
			)
			continue
		}

		if !strings.Contains(url, "http") {

			codeBlock := `
- *Example* *:* %s
.play %s/%s
			`
			document = strings.ReplaceAll(
				document,
				m,
				fmt.Sprintf(
					codeBlock,
					title,
					modName,
					url,
				),
			)

			continue
		}

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

	//fmt.Println(document)
	err = os.WriteFile(
		filePath,
		[]byte(document),
		os.ModePerm,
	)

	if err != nil {
		panic(err)
	}

}
