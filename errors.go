package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleError(message string, errorCode int, w http.ResponseWriter) {
	type ErrorReturnVals struct {
		Error string `json:"error"`
	}
	respBody := ErrorReturnVals{
		Error: message,
	}
	data, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	w.Write(data)
}
