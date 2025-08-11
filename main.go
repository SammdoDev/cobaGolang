package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go API on Render!")
	})

	port := os.Getenv("PORT") // Render akan kasih PORT di environment variable
	if port == "" {
		port = "8080" // default untuk lokal
	}

	http.ListenAndServe(":"+port, nil)
}
