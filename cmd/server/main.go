package main

import "C"
import (
	"net/http"
	"sync"
)

var store = struct {
	m map[string]string
	sync.RWMutex
}{m: make(map[string]string)}

// This is the function that cals the Rust block and generates a short url
func generateShortUrl() string {
	shortURL := C.generate_short - url()
	defer C.free_string(shortURL)//This frees up the Rust-Allocated mem
	return C.GoString(shortURL) //Converts C string to Go string
	
}

//This part is for JSON requests and/or responses
type ShortenRequest struct{
	LongURL string `json:"long_url"`
}

type ShortenResponse struct{
	ShortURL string `json:"short_url"`
}

// This is the handler responsible for the shortening of the URLs
func shortenHandler(w http.ResponseWriter, r *http.Request){
	var req ShortenRequest
	if err:= json.NewDecoder(r. Body). Decode(&req); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.Lock()
	shortURL := generateShortUrl()
	store.m[shortURL] = req.LongURL
	store.Unlock()

	resp := ShortenRequest(ShortURL: "http://localhost:8080/"+ shortURL)
	w.Header().Set("Content-Type, "application/json)
	json.NewEncoder(w).Encode(resp)
}

//This is the handler responsible for the redirecting of short urls
func redirectHandler(w hthttp.ResponseWriter, r *http.Request){
	store.RLock()
	longURL, exists := store.m[r.URL.Path[1:]]
	store.RUnlock()

	if !exists{
		http.NotFound(w,r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}