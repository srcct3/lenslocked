package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTmpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(filepath string) (Template, error) {

	t, err := template.ParseFiles(filepath)

	if err != nil {
		return Template{}, fmt.Errorf("Parsing template: %w", err)
	}

	return Template{
		htmlTmpl: t,
	}, nil
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := t.htmlTmpl.Execute(w, data)

	if err != nil {
		log.Printf("Error excuting template: %v", err)
		http.Error(w, "Error excuting template", http.StatusInternalServerError)
	}
}
