package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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
