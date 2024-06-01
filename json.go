package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 { //Server error status codes
		log.Printf("Responding with %d: %s", code, message)
	}
	type errResponse struct {
		Error string `json."error"`
	}
	responseWithJSON(w, code, errResponse{
		Error: message,
	})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	if data, err := json.Marshal(payload); err != nil {
		log.Printf("Failed to write JSON: %v\n", err)
		w.WriteHeader(500)
	} else {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(code)
		w.Write(data)
	}
}
