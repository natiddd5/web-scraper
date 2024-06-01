package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {

	//convert GO data into JSON format.
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON response - ", payload)
		w.WriteHeader(500)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func errorResponse(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 500 level error", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	jsonResponse(w, code, errorResponse{message})
}
