package router

import (
	"fmt"
	"net/http"
)

// ヘルスチェック
func handleHealth(mux *http.ServeMux) {
	mux.HandleFunc("GET /health/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "health check")
	})
}
