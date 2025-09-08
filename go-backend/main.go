package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

func getRandomQuote() Quote {
	quotes := []Quote{
		{"Stay hungry, stay foolish.", "Steve Jobs"},
		{"Simplicity is the ultimate sophistication.", "Leonardo da Vinci"},
		{"Code is like humor. When you have to explain it, it’s bad.", "Cory House"},
	}
	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getRandomQuote())
}
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3002")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/quote", quoteHandler)

    http.ListenAndServe(":8080", enableCORS(mux))
}
