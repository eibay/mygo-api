package main

import (
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

type ApiHandler struct{}

func (ah *ApiHandler) handleHello(w http.ResponseWriter, r *http.Request) {
	response := "Hello World"

	w.Write([]byte(response))
}

func main() {
	apiHandler := ApiHandler{}
	http.HandleFunc("/hello", apiHandler.handleHello)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
