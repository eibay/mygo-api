package main

import "net/http"

type ApiHandler struct{}

func (ah *ApiHandler) handleHello(w http.ResponseWriter, r *http.Request) {
	response := "Hello World"

	w.Write([]byte(response))
}
