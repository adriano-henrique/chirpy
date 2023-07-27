package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type metricsApiConfig struct {
	fileServerHits int
}

func apiRouter(configMetrics *metricsApiConfig) http.Handler {
	r := chi.NewRouter()
	r.Get("/healthz", handlerReadiness)
	r.Get("/metrics", configMetrics.handlerMetrics)
	return r
}

func adminRouter(configMetrics *metricsApiConfig) http.Handler {
	r := chi.NewRouter()
	r.Get("/metrics", configMetrics.handlerAdminMetrics)
	return r
}

func main() {
	const port = "8080"
	const filePathRoot = "."

	configMetrics := metricsApiConfig{
		fileServerHits: 0,
	}

	r := chi.NewRouter()
	fsHandler := configMetrics.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot))))
	r.Handle("/app/*", fsHandler)
	r.Handle("/app", fsHandler)
	r.Mount("/api", apiRouter(&configMetrics))
	r.Mount("/admin", adminRouter(&configMetrics))

	corsMux := middlewareCors(r)
	newServer := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filePathRoot, port)
	log.Fatal(newServer.ListenAndServe())
}
