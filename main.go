package main

import (
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8000"
)

func main() {
	apiHandler := ApiHandler{}
	http.HandleFunc("/hello", apiHandler.handleHello)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
