// markdown-to-site utility to convert
// markdown to format used by the gotour
// project. Support exists for links, bullet-point lists
// and level 3 headers denoted with '###'.
// Run with flag debug to print the text without overwriting the target
// file.
// The program will overwrite the specified file_path. The contents
// of file_path should be markdown on first run.
// Markdown URLs with parentheses will break the program.
// for example, [Sample (Stuff)](http...)
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
	var debug bool

	flag.StringVar(
		&filePath,
		"file_path",
		"./_content/tour/pointers.article",
		"Path to Markdown file",
	)

	flag.BoolVar(
		&debug,
		"debug",
		false,
		"Print generated file instead of writing.",
	)

	flag.StringVar(
		&modName,
		"module_name",
		"pointers",
		"Name of module",
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

	if debug {
		fmt.Println(document)
		return
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
