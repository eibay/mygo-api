package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {
	apiHandler := ApiHandler{}

	http.HandleFunc("/hello", apiHandler.handleHello)
	http.HandleFunc("/health", apiHandler.handleHealthCheck)
	http.HandleFunc("/metadata", apiHandler.handleMetadata)

	fmt.Println("Server started and listening at port:", port)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
