package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	authHandler "github.com/kakkky/app/adapter/presentation/user/auth"
	"github.com/kakkky/app/adapter/repository"
	authUsecase "github.com/kakkky/app/application/usecase/user/auth"
	"github.com/kakkky/app/infrastructure/auth"
)

func handleAuth(mux *http.ServeMux) {
	authorization := middleware.Authorication(
		auth.NewJWTAuthenticator(),
		repository.NewTokenAuthenticatorRepository(),
	)

	mux.Handle("POST /login", composeMiddlewares(middleware.Logger)(
		authHandler.NewLoginHandler(
			authUsecase.NewLoginUsecase(
				repository.NewUserRepository(),
				repository.NewTokenAuthenticatorRepository(),
				auth.NewJWTAuthenticator(),
			))))

	mux.Handle("DELETE /logout", composeMiddlewares(authorization, middleware.Logger)(
		authHandler.NewLogoutHandler(
			authUsecase.NewLogoutUsecase(
				auth.NewJWTAuthenticator(),
				repository.NewTokenAuthenticatorRepository(),
			),
		),
	))

}
