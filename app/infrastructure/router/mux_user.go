package router

import (
	"net/http"

	userHandler "github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/adapter/repository"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/user"
)

func handleUser(mux *http.ServeMux) {
	userRepository := repository.NewUserRepository()

	mux.Handle("POST /user", userHandler.NewPostUserHandler(userUsecase.NewRegisterUsecase(
		userRepository, user.NewUserDomainService(userRepository),
	)))
	mux.Handle("DELETE /user/{id}", userHandler.NewDeleteUserHandler(userUsecase.NewUnregisterUsecase(
		userRepository,
	)))
	mux.Handle("GET /users", userHandler.NewGetUsersHandler(userUsecase.NewListUsersUsecase(
		userRepository,
	)))
}
