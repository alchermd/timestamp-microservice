package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const API_URL = "/api/"
const TIMESTAMP_URL = "timestamp/"
const INDEX_URL = "/"

type Timestamp struct {
	Unix int64  `json:"unix"`
	UTC  string `json:"utc"`
}

type Message struct {
	Message string `json:"message"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, nil)
}

func timestampHandler(w http.ResponseWriter, r *http.Request) {
	dateString := r.URL.Path[len(API_URL+TIMESTAMP_URL):]

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	t, err := time.Parse("2006-01-02", dateString)

	if err != nil {
		j, _ := json.Marshal(&Message{Message: "Invalid Date"})
		fmt.Fprint(w, string(j))
		return
	}

	ts := &Timestamp{
		Unix: t.Unix(),
		UTC:  t.UTC().String(),
	}

	j, _ := json.Marshal(ts)
	fmt.Fprint(w, string(j))
}

func main() {
	http.HandleFunc(API_URL+TIMESTAMP_URL, timestampHandler)
	http.HandleFunc(INDEX_URL, indexHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
