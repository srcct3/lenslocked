package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func excuteTemplate(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles(filepath)

	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Printf("Error excuting template: %v", err)
		http.Error(w, "Error excuting template", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "home.tmpl")
	excuteTemplate(w, path)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "contact.tmpl")
	excuteTemplate(w, path)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("templates", "faq.tmpl")
	excuteTemplate(w, path)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Server at :8080")
	http.ListenAndServe(":8080", r)
}
