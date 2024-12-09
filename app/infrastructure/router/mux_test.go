package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantCode int
		wantErr  bool
	}{
		{
			name:     "正常にヘルスチェックハンドラを呼び出せる",
			method:   http.MethodGet,
			path:     "/health",
			wantCode: http.StatusOK,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.path, nil)
			sut := NewMux()
			sut.ServeHTTP(w, r)
			resp := w.Result()
			if resp.StatusCode != tt.wantCode {
				t.Errorf("want status code %d, but got %d", tt.wantCode, resp.StatusCode)
			}
		})
	}
}

func TestComposeMiddlewares(t *testing.T) {
	middleware1 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Middleware", "1")
			h.ServeHTTP(w, r)
		})
	}
	middleware2 := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Middleware", "2")
			h.ServeHTTP(w, r)
		})
	}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	sut := composeMiddlewares(middleware1, middleware2)(handler)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	sut.ServeHTTP(w, r)
	// 最終的なMiddrewareヘッダーは2である。
	// リクエスト処理は、Middleware1→Middleware2と行われる
	if got := w.Header().Get("Middleware"); got != "2" {
		t.Errorf("expected header Middleware to be '2', got '%s'", got)
	}
}
