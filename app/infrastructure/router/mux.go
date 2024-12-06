package router

import (
	"net/http"
)

// ルーティングを登録したマルチプレクサを返す
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	handleHealth(mux)
	return mux
}

// 適用させたい順で、ミドルウェアを引数に入れる
// composeMiddewares(M1,M2,M3)とした場合、M1(M2(M3()))といったようにラップされたミドルウェアを返す
func composeMiddlewares(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-(i+1)](h)
		}
		return h
	}
}
