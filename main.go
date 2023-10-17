package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/eliasyeme/lenslocked/controllers"
	"github.com/eliasyeme/lenslocked/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	t, err := views.Parse(filepath.Join("templates", "home.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(t))

	t, err = views.Parse(filepath.Join("templates", "contact.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(t))

	t, err = views.Parse(filepath.Join("templates", "faq.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(t))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Server at :8080")
	http.ListenAndServe(":8080", r)
}
