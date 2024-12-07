package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/health"
)

func handleHealth(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", health.HealthCheckHandler)
}
