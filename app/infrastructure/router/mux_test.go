package router

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewMux(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		want     string
		wantCode int
		wantErr  bool
	}{
		{
			name:     "正常にヘルスチェックハンドラを呼び出せる",
			method:   http.MethodGet,
			path:     "/health/",
			want:     "health check",
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
			got, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("failed to read response body : %v", err)
			}
			if string(got) != tt.want {
				t.Errorf("want %q,but got %q", tt.want, got)
			}
		})
	}

}

func TestComposeMiddlewares(t *testing.T) {
	middleware1 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Middleware", "1")
			next.ServeHTTP(w, r)
		})
	}

	middleware2 := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Middleware", "2")
			next.ServeHTTP(w, r)
		})
	}

	// Define a final handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Compose middlewares
	composed := composeMiddlewares(middleware1, middleware2)(handler)

	// Create a test request and response
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Serve the request
	composed.ServeHTTP(rec, req)

	// 最終的なMiddrewareヘッダーは2である。
	if got := rec.Header().Get("Middleware"); got != "2" {
		t.Errorf("Expected header Middleware to be '2', got '%s'", got)
	}
}
