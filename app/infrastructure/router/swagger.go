package router

import (
	"net/http"

	swagger "github.com/swaggo/http-swagger"
)

func handleSwagger(mux *http.ServeMux) {
	mux.Handle("GET /swagger/", swagger.WrapHandler)
}
