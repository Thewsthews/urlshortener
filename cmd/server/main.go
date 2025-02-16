package main

/*
#include <stdlib.h>
#include <string.h>

char* generate_short_url() {
	const char* short_url = "shortURL123"; // Example short URL
	char* result = (char*)malloc(strlen(short_url) + 1);
	strcpy(result, short_url);
	return result;
}
*/
import "C"
import (
	"encoding/json"
	"net/http"
	"sync"
)

var store = struct {
	m map[string]string
	sync.RWMutex
}{m: make(map[string]string)}

// This is the function that cals the Rust block and generates a short url
func generateShortUrl() string {
	shortURL := C.generate_short_url()
	return C.GoString(shortURL) // Converts C string to Go string
	//Converts C string to Go string

}

// This part is for JSON requests and/or responses
type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// This is the handler responsible for the shortening of the URLs
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Lock()
	shortURL := generateShortUrl()
	store.m[shortURL] = req.LongURL
	store.Unlock()

	resp := ShortenResponse{ShortURL: "http://localhost:8080/" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// This is the handler responsible for the redirecting of short urls
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	store.RLock()
	longURL, exists := store.m[r.URL.Path[1:]]
	store.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}
func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
