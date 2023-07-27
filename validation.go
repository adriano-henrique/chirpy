package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handlerValidation(w http.ResponseWriter, req *http.Request) {
	type Parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(req.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		handleError("Something went wrong", http.StatusInternalServerError, w)
		return
	}

	if len(params.Body) > 140 {
		handleError("Chirp is too long", 400, w)
		return
	}

	type ValidReturnVals struct {
		Valid bool `json:"valid"`
	}
	respBody := ValidReturnVals{
		Valid: true,
	}
	response, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
