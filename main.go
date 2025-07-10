package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlContent, err := os.ReadFile("index.html")
	if err != nil {
		// If HTML file doesn't exist, serve a fallback message
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "HTML file not found. Error:", err)
		return
	}
	
	// Set content type to HTML
	w.Header().Set("Content-Type", "text/html")
	
	// Write the HTML content
	w.Write(htmlContent)
}

func main() {
	// Get port from environment variable (for App Engine) or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	http.HandleFunc("/", handler)
	fmt.Printf("Server running on port %s...\n", port)
	fmt.Printf("Visit http://localhost:%s to see your HTML page!\n", port)
	http.ListenAndServe(":"+port, nil)
}
