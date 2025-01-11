package router

import (
	"net/http"

	userHandler "github.com/kakkky/app/adapter/presentation/handler/user"
	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/application/usecase/auth"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/user"
	authInfra "github.com/kakkky/app/infrastructure/auth"
	"github.com/kakkky/app/infrastructure/db/sqlc"
	"github.com/kakkky/app/infrastructure/kvs"
)

func handleUser(mux *http.ServeMux) {
	userRepository := repository.NewUserRepository(sqlc.NewSqlcQuerier())
	authorization := middleware.Authorication(
		auth.NewAuthorizationUsecase(
			authInfra.NewJWTAuthenticator(),
			repository.NewTokenAuthenticatorRepository(kvs.NewRedisCommander()),
		),
	)

	mux.Handle("POST /user", composeMiddlewares(middleware.Logger)(
		userHandler.NewPostUserHandler(
			userUsecase.NewRegisterUsecase(
				userRepository,
				user.NewUserDomainService(userRepository),
			))))
	mux.Handle("DELETE /user", composeMiddlewares(authorization, middleware.Logger)(
		userHandler.NewDeleteUserHandler(
			userUsecase.NewUnregisterUsecase(
				userRepository,
			))))
	mux.Handle("GET /users", composeMiddlewares(authorization, middleware.Logger)(
		userHandler.NewGetUsersHandler(
			userUsecase.NewFetchUsersUsecase(
				userRepository,
			))))
	mux.Handle("PATCH /user", composeMiddlewares(authorization, middleware.Logger)(
		userHandler.NewUpdateUserHandler(userUsecase.NewUpdateProfileUsecase(
			userRepository,
		))))
}
