package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the URL Shortener API!")
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Mthod Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var newURL URLRequest
	err := json.NewDecoder(r.Body).Decode(&newURL)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	shortCode := generateShortCode()
	urlStore[shortCode] = newURL.URL
	response := map[string]string{
		"short_url": "http://localhost:8080/" + shortCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		homeHandler(w, r)
		return
	}

	if r.URL.Path == "/shorten" {
		shortenURLHandler(w, r)
		return
	}

	shortCode := r.URL.Path[1:]

	originalURL, exists := urlStore[shortCode]
	if !exists {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
