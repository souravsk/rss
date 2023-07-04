package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Responce with 500 error ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, status, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(status)
	w.Write(data)

}
