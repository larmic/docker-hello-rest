package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greet struct {
	Text string `json:"Greet"`
}

func serveHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")

	greet := Greet{Text: "Hello (REST-) World!"}

	_ = json.NewEncoder(w).Encode(greet)
}

func main() {
	log.Println("Hello docker_hello_rest")

	http.HandleFunc("/hello", serveHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
