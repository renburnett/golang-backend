package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getAPIKey() string {
	file, err1 := os.Open("API_KEY.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	val, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		log.Fatal(err2)
	}
	return string(val)
}

func main() {
	fetchGPSData(getAPIKey())
	handleRequests()
}

func fetchGPSData(apiKey string) {
	uri := "https://track.onestepgps.com/v3/api/public/device?latest_point=true&api-key=" + apiKey
	response, err1 := http.Get(uri)

	if err1 != nil {
		// handle error
		print(err1)
	}
	defer response.Body.Close()

	var gpsList map[string]interface{}

	err2 := json.NewDecoder(response.Body).Decode(&gpsList)
	if err2 != nil {
		print(err2)
	}
	fmt.Print("response", gpsList)
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
