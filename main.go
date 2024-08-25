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

func SimpleConvert(word string) string {
	new := word[:3]
	return new
}

func FirstEndConvert(word string) string {
	new := string(word[0]) + string(word[len(word)-1])
	return new
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

	abbreviation := Abbreviation{Abbreviation: SimpleConvert(word.Word)}

	json.NewEncoder(w).Encode(abbreviation)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/convert", convert)

	log.Printf("Serving on port 3333 ...")
	log.Fatal(http.ListenAndServe(":3333", nil))
}
