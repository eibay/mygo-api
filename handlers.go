package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiHandler struct{}

type ResponseHealthCheck struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

func (ah *ApiHandler) handleHello(w http.ResponseWriter, r *http.Request) {
	response := "Hello World"

	w.Write([]byte(response))
}

func (ah *ApiHandler) handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	responseHealthCheck := ResponseHealthCheck{
		Status:      "Ok",
		Description: "Basic API Service",
	}

	response, err := json.Marshal(&responseHealthCheck)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}
