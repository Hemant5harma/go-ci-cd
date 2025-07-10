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
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080...")
	fmt.Println("Visit http://localhost:8080 to see your HTML page!")
	http.ListenAndServe(":8080", nil)
}
