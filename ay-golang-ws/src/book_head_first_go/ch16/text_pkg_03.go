package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	templateText := "{{if.}}这是{{end}}一个"
	tmpl, err := template.New("test").Parse(templateText)
	tmpl.Execute(os.Stdout, false)

	if err != nil {
		log.Fatal(err)
	}
}
