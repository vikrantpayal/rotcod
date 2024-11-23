package main

import (
	"html/template"
	"net/http"
)

// Struct to pass data to templates
type PageData struct {
	PageName string
}

func renderPage(w http.ResponseWriter, r *http.Request, pageTitle, templateFile string) {
	// Prepare data with the dynamic page title
	data := PageData{PageName: pageTitle}

	// Parse the HTML template
	tmpl := template.Must(template.ParseFiles(templateFile, "header.html"))
	tmpl.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "Home", "home.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, r, "About Us", "about.html")
}

func main() {
	// Serve static files (CSS, images, JS, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handlers for pages
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Launch the server
	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
