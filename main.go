package main

import (
	"fmt"
	"log"
	"net/http"
)

type metricsApiConfig struct {
	fileServerHits int
}

func (cfg *metricsApiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits++
		next.ServeHTTP(w, r)
	})
}

func (cfg *metricsApiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileServerHits)))
}

func main() {
	const port = "8080"
	const filePathRoot = "."

	configMetrics := metricsApiConfig{
		fileServerHits: 0,
	}

	r := chi.NewRouter()

	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", configMetrics.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot)))))
	serveMux.HandleFunc("/healthz", handlerReadiness)
	serveMux.HandleFunc("/metrics", configMetrics.handlerMetrics)

	corsMux := middlewareCors(serveMux)
	newServer := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filePathRoot, port)
	log.Fatal(newServer.ListenAndServe())
}
