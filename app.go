package main

import (
	"log"
	"net/http"
	"html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, nil)
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func main() {
	const API_URL = "/api/"
	const TIMESTAMP_URL = "timestamp/"
	const INDEX_URL = "/"
	
	http.HandleFunc(API_URL + TIMESTAMP_URL, timestampHandler)
	http.HandleFunc(INDEX_URL, indexHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}