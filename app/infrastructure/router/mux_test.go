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
			rw := httptest.NewRequest(tt.method, tt.path, nil)
			sut := NewMux()
			sut.ServeHTTP(w, rw)
			resp := w.Result()
			if resp.StatusCode != tt.wantCode {
				t.Errorf("want status code %d, but got %d", tt.wantCode, resp.StatusCode)
			}
		})
	}
}
