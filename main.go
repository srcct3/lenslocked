package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Home page</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>Contact page</h1>
		<p><a href="mailto:me@you.com">Contact us</a></p>
	`)
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

type Router struct{}

func (route Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router

	fmt.Println("Server at :8080")
	http.ListenAndServe(":8080", router)
}
