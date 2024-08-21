package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Word struct {
	Word string `json:"word"`
}

type Abbreviation struct {
	Abbreviation string `json:"abbreviation"`
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func convert(w http.ResponseWriter, r *http.Request) {

	var word Word

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	abbreviation := Abbreviation{Abbreviation: "Not Implemented!"}

	json.NewEncoder(w).Encode(abbreviation)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/convert", convert)

	log.Printf("Serving on port 3333 ...")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
