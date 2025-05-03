package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// code -> the status code we are gonna respond with.
// payload -> something which can be "marshelled" JSON structure.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marchal the payload: %v\n", payload)
		log.Printf("Marshal error %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
