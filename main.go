package main

import (
	"net/http"
	"time"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.Handle("/", http.FileServer(http.Dir(".")))

	corsMux := middlewareCors(serverMux)
	newServer := &http.Server{
		Addr:           ":8080",
		Handler:        corsMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	newServer.ListenAndServe()
}
