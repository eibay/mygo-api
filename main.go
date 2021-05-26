package main

import (
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8001"
)

func main() {
	apiHandler := ApiHandler{}

	http.HandleFunc("/hello", apiHandler.handleHello)
	http.HandleFunc("/health", apiHandler.handleHealthCheck)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
