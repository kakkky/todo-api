package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	userHandler "github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/adapter/repository"
	"github.com/kakkky/app/application/usecase/auth"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/user"
	authInfra "github.com/kakkky/app/infrastructure/auth"
)

func handleUser(mux *http.ServeMux) {
	userRepository := repository.NewUserRepository()
	authorization := middleware.Authorication(
		auth.NewAuthorizationUsecase(
			authInfra.NewJWTAuthenticator(),
			repository.NewTokenAuthenticatorRepository(),
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
