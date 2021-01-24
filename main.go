package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", landingPage)
	http.HandleFunc("/echo", echo) //TODO remove
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Article struct
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles struct array
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "ComputerRAM", Desc: "New RAM technologies", Content: "Yooooooooooo this ram is fast"},
	}
	fmt.Println("Articles endpoint hit!")
	json.NewEncoder(w).Encode(articles)
}

func landingPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Landing Page Hit")
}

func echo(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Query()["message"][0]
	fmt.Println(message)
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}
