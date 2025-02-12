package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const shortURLLength = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var store = struct {
	m map[string]string
	sync.RWMutex
}{m: make(map[string]string)}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, shortURLLength)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Lock()
	shortURL := generateShortURL()
	store.m[shortURL] = req.LongURL
	store.Unlock()

	resp := ShortenResponse{ShortURL: "http://localhost:8080/shorten" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

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
