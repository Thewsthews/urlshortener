package main

import (
	"math/rand"
	"time"
)

const shortURLLength = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var store = struct {
	m map[string]string
	sunc.RWMutex
}{m: make(map[string]string)}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	shortURL
}
