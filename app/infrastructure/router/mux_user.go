package router

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	userHandler "github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/adapter/repository"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/user"
	"github.com/kakkky/app/infrastructure/auth"
)

func handleUser(mux *http.ServeMux) {
	userRepository := repository.NewUserRepository()
	authorization := middleware.AuthoricationController(auth.NewJWTAuthenticator(), repository.NewTokenAuthenticatorRepository())

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
			userUsecase.NewListUsersUsecase(
				userRepository,
			))))
	mux.Handle("PATCH /user", composeMiddlewares(middleware.Logger)(
		userHandler.NewUpdateUserHandler(userUsecase.NewUpdateProfileUsecase(
			userRepository,
		))))
}
