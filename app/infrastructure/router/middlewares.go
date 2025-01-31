package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	authUsecase "github.com/kakkky/app/application/usecase/auth"
)

var (
	authorization func(h http.Handler) http.Handler
	logger        func(h http.Handler) http.Handler
)

// ミドルウェアを初期化する
func initMiddlewares() {
	authorization = middleware.Authorication(
		authUsecase.NewAuthorizationUsecase(
			jwtAuthenticator,
			jwtAuthenticatorRepository,
		),
	)
	logger = middleware.Logger
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
