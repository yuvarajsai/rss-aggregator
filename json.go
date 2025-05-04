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
		log.Printf("Failed to marshal the payload: %v\n", payload)
		log.Printf("Marshal error %v\n", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
		// In the below structure, only the fields which start with an uppercase letter will be exported and marshalled using encoding/json.
		// `json:"errorCode"` etc. should be defined, so that the text will be serialized into keys and the values will be stored in the keys.
		type errorResponse struct {
			ErrorCode    int    `json:"errorCode"`
			ErrorMessage string `json:"errorMessage"`
		}
		errorResp := errorResponse{
			ErrorCode:    code,
			ErrorMessage: msg,
		}
		respondWithJSON(w, code, errorResp)
	}
}
