package main

import (
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
}
