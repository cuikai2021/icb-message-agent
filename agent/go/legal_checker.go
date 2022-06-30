package agent

import (
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed ../../templates/*
var templates embed.FS

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
	content, _ := templates.ReadFile("messages.txt")
	fmt.Println(content)
	//files, err := ioutil.ReadDir(templatePath)
	//if err != nil {
	//	return err
	//}
	//
	//for _, file := range files {
	//	inFile, err := os.Open(filepath.Join(templatePath, file.Name()))
	//	if err != nil {
	//		return err
	//	}
	//	defer inFile.Close()
	//
	//	scanner := bufio.NewScanner(inFile)
	//	for scanner.Scan() {
	//		c.msgTemplates[scanner.Text()] = true
	//	}
	//}

	return nil
}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
