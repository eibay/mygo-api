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

type ResponseMetadata struct {
	Version       string `json:"version"`
	Description   string `json:"description"`
	LastCommitSha string `json:"lastcommitsha"`
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

func (ah *ApiHandler) handleMetadata(w http.ResponseWriter, r *http.Request) {
	responseMetaData := ResponseMetadata{
		Version:       "1.0.7",
		Description:   "basic go-lang api service",
		LastCommitSha: "oasdfkjqoewrpsdfklnvkjowo",
	}

	response, err := json.Marshal(&responseMetaData)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(response)
}
