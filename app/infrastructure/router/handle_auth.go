package router

import (
	"net/http"

	authHandler "github.com/kakkky/app/adapter/presentation/handler/auth"
	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/repository"
	authUsecase "github.com/kakkky/app/application/usecase/auth"
	authInfra "github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
)

func handleAuth(mux *http.ServeMux) {
	authorization := middleware.Authorication(
		authUsecase.NewAuthorizationUsecase(
			authInfra.NewJWTAuthenticator(),
			repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander()),
		),
	)

	mux.Handle("POST /login", composeMiddlewares(middleware.Logger)(
		authHandler.NewLoginHandler(
			authUsecase.NewLoginUsecase(
				repository.NewUserRepository(sqlc.NewSqlcQuerier()),
				repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander()),
				authInfra.NewJWTAuthenticator(),
			))))

	mux.Handle("DELETE /logout", composeMiddlewares(authorization, middleware.Logger)(
		authHandler.NewLogoutHandler(
			authUsecase.NewLogoutUsecase(
				authInfra.NewJWTAuthenticator(),
				repository.NewJwtAuthenticatorRepository(kvs.NewRedisCommander()),
			),
		),
	))

}
