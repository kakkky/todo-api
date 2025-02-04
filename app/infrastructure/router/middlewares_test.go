package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestComposeMiddlewares(t *testing.T) {
	middleware1 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Middleware", "1")
			h.ServeHTTP(rw, r)
		})
	}
	middleware2 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Middleware", "2")
			h.ServeHTTP(rw, r)
		})
	}
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {})

	sut := composeMiddlewares(middleware1, middleware2)(handler)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rw := httptest.NewRecorder()
	sut.ServeHTTP(rw, r)
	// 最終的なMiddrewareヘッダーは2である。
	// リクエスト処理は、Middleware1→Middleware2と行われる
	if got := rw.Header().Get("Middleware"); got != "2" {
		t.Errorf("expected header Middleware to be '2', got '%s'", got)
	}
}
