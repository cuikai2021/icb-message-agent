package agent

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

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
	templatePath, _ := filepath.Abs("../../templates")
	fmt.Println(templatePath)

	files, err := ioutil.ReadDir(templatePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		inFile, err := os.Open(filepath.Join(templatePath, file.Name()))
		if err != nil {
			return err
		}
		defer inFile.Close()

		scanner := bufio.NewScanner(inFile)
		for scanner.Scan() {
			c.msgTemplates[scanner.Text()] = true
		}
	}

	return nil
}
