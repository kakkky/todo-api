package router

import (
	"net/http"

	authHandler "github.com/kakkky/app/adapter/presentation/auth"
	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/application/usecase/auth"
	authUsecase "github.com/kakkky/app/application/usecase/auth"
	authInfra "github.com/kakkky/app/infrastructure/auth"
)

func handleAuth(mux *http.ServeMux) {
	authorization := middleware.Authorication(
		auth.NewAuthorizationUsecase(
			authInfra.NewJWTAuthenticator(),
			repository.NewTokenAuthenticatorRepository(),
		),
	)

	mux.Handle("POST /login", composeMiddlewares(middleware.Logger)(
		authHandler.NewLoginHandler(
			authUsecase.NewLoginUsecase(
				repository.NewUserRepository(),
				repository.NewTokenAuthenticatorRepository(),
				authInfra.NewJWTAuthenticator(),
			))))

	mux.Handle("DELETE /logout", composeMiddlewares(authorization, middleware.Logger)(
		authHandler.NewLogoutHandler(
			authUsecase.NewLogoutUsecase(
				authInfra.NewJWTAuthenticator(),
				repository.NewTokenAuthenticatorRepository(),
			),
		),
	))

}
