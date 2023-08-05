package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

type TemplateVariables struct {
	Name string
}

var funcMap = template.FuncMap{
	"CustomFunc": func(dictionary string) string {
		return "return " + dictionary
	},
	"Channels": func() []string {
		ch := make(chan string)
		done := make(chan bool)
		go func() {
			for i := 0; i < 10; i++ {
				ch <- "str" + strconv.Itoa(i)
			}
			close(ch)
			done <- true
		}()
		result := make([]string, 0, 10)

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case str := <-ch:
				result = append(result, str)
				break LOOP
			}
		}
		return result
	},
}

func main() {
	if err := runMain(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func templateFromString(templateVars TemplateVariables, funcMap template.FuncMap) error {
	templateString := `{{ .Name }}
{{ CustomFunc "argument" }}`
	tmpl := template.Must(template.New("template").Funcs(funcMap).Parse(templateString))
	err := tmpl.Execute(os.Stdout, templateVars)
	if err != nil {
		return err
	}
	return nil
}

func templateFromFile(templateVars TemplateVariables, funcMap template.FuncMap) error {
	baseFilePath := "./base.go.tmpl"
	entryFilePath := "./main.go.tmpl"
	baseFileName := filepath.Base(baseFilePath)
	tmpl := template.Must(template.New(baseFileName).Funcs(funcMap).ParseFiles(baseFilePath, entryFilePath))
	return tmpl.Execute(os.Stdout, templateVars)
}

func runMain() error {
	vars := TemplateVariables{
		Name: "name inserted",
	}
	// return templateFromString(vars, funcMap)
	return templateFromFile(vars, funcMap)
}
