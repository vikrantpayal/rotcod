package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the static files like CSS, JS, and the HTML file
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
