package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Respond with 5xx error %v:", msg)
	}

	type errorResponse struct {
		Error string
	}

	respondWithJSON(w, code, errorResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json respone %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
