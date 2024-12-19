package auth

import (
	"encoding/json"
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user/auth"
	"github.com/kakkky/pkg/validation"
)

type LoginHandler struct {
	loginUsecase *auth.LoginUsecase
}

func NewLoginHandler(loginUsecase *auth.LoginUsecase) *LoginHandler {
	return &LoginHandler{
		loginUsecase: loginUsecase,
	}
}

func (lh *LoginHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var params LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidation().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	input := auth.LoginUsecaseInputDTO{
		Email:    params.Email,
		Password: params.Password,
	}
	output, err := lh.loginUsecase.Run(r.Context(), input)
	if err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	resp := LoginResponse{
		SignedToken: output.SignedToken,
	}
	presenter.RespondOK(rw, resp)
}
