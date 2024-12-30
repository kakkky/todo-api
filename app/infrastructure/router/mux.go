package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/handler/health"
	swagger "github.com/swaggo/http-swagger"
)

// ルーティングを登録したマルチプレクサを返す
func NewMux() http.Handler {
	mux := http.NewServeMux()
	// 開発用ルーティング
	{
		mux.HandleFunc("GET /health", health.HealthCheckHandler)
		mux.Handle("GET /swagger/", swagger.WrapHandler)
	}
	// ユーザー系
	{
		handleUser(mux)
		handleAuth(mux)
	}
	// タスク系
	{
		handleTask(mux)
	}

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
