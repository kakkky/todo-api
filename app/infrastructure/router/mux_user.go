package router

import (
	"net/http"

	userHandler "github.com/kakkky/app/adapter/presentation/user"
	"github.com/kakkky/app/adapter/repository"
	userUsecase "github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/user"
)

func handleUser(mux *http.ServeMux) {
	mux.Handle("POST /user", userHandler.NewPostUserHandler(userUsecase.NewRegisterUsecase(
		repository.NewUserRepository(), user.NewUserDomainService(repository.NewUserRepository()),
	)))
}
