package agent

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"path"
	"strings"
)

//go:embed icb-message-templates/*.txt
var fs embed.FS

type legalChecker struct {
	msgTemplates map[string]bool
}

func NewLegalChecker() *legalChecker {
	checker := &legalChecker{
		msgTemplates: make(map[string]bool),
	}

	if err := checker.loadMsgTemplates(); err != nil {
		log.Print(err)
	}

	return checker
}

func (c *legalChecker) IsLegal(message string) bool {
	if c.msgTemplates[message] != true {
		return false
	} else {
		return true
	}

}

func (c *legalChecker) loadMsgTemplates() (err error) {
	filePaths, err := getAllFilenames(&fs, "icb-message-templates")

	for _, filePath := range filePaths {
		content, _ := fs.ReadFile(filePath)
		templateLines := SplitLines(string(content))
		for _, templateLine := range templateLines {
			c.msgTemplates[templateLine] = true
		}
	}

	fmt.Printf("%v", c.msgTemplates)

	return nil
}

func SplitLines(content string) []string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(content))
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func getAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}

	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}

			out = append(out, res...)

			continue
		}

		out = append(out, fp)
	}

	return
}
