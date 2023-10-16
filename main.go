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
	filepath := filepath.Join("templates", "home.tmpl")
	excuteTemplate(w, filepath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	filepath := filepath.Join("templates", "contact.tmpl")
	excuteTemplate(w, filepath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	faq := map[string]string{
		"whats your name":         "my name is elias",
		"how old are you":         "I am 28 years old",
		"do you have any sibling": "I have just one brother",
	}
	fmt.Fprint(w, `
		<h1>FAQ</h1>
	`)
	for k, v := range faq {
		fmt.Fprint(w, fmt.Sprintf(`
		<h2>%v</h2>
		<p>%v</p>
	`, k, v))
	}
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
