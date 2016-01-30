package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port = flag.String("port", env("PORT", "8080"), "Listener port")
)

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

type Payload struct {
	Method  string `json:"method"`
	Request string `json:"request"`
	Time    string `json:"time"`
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] - %s%s", r.Method, r.Host, r.URL)

	payload := Payload{r.Method, fmt.Sprintf("%s%s", r.Host, r.URL), time.Now().String()}

	json, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	flag.Parse()

	log.Printf("Starting server on port %s", *port)

	http.HandleFunc("/", handle)
	http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
}
