package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

func (cfg *metricsApiConfig) handlerAdminMetrics(w http.ResponseWriter, req *http.Request) {
	content, error := ioutil.ReadFile("admin.html")
	if error != nil {
		log.Fatal(error)
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(string(content), cfg.fileServerHits)))
}
