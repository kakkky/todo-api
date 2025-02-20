package middleware

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
)

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 開発環境ようなので全てのリクエストを通すようにする
		w.Header().Set("Access-Control-Allow-Origin", "*") // 本番環境でこのようにするのは禁止
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Requested-With, Origin, X-Csrftoken, Accept, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PATCH, DELETE, OPTIONS")
		// プリフライトリクエストの対応
		if r.Method == "OPTIONS" {
			presenter.RespondNoContent(w)
			return
		}
		h.ServeHTTP(w, r)
	})
}
