package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dateString := r.URL.Path[len(API_URL+TIMESTAMP_URL):]

	ts, err := parseDateString(dateString)

	if err != nil {
		j, _ := json.Marshal(&Message{Message: "Invalid Date"})
		fmt.Fprint(w, string(j))
		return
	}

	j, _ := json.Marshal(ts)
	fmt.Fprint(w, string(j))
}

func parseDateString(dateString string) (*Timestamp, error) {
	var t time.Time
	var err error

	if dateString == "" {
		return &Timestamp{
			Unix: time.Now().Unix(),
			UTC:  time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 MST"),
		}, nil
	}

	if !strings.Contains(dateString, "-") {
		var i int64
		i, err = strconv.ParseInt(dateString, 10, 64)

		if err != nil {
			return nil, err
		}

		t = time.Unix(i, 0)
	} else {
		t, err = time.Parse("2006-01-02", dateString)

		if err != nil {
			return nil, err
		}
	}

	ts := &Timestamp{
		Unix: t.Unix(),
		UTC:  t.UTC().Format("Mon, 2 Jan 2006 15:04:05 MST"),
	}

	return ts, nil
}

func main() {
	http.HandleFunc(API_URL+TIMESTAMP_URL, timestampHandler)
	http.HandleFunc(INDEX_URL, indexHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
