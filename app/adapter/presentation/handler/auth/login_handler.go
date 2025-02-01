package auth

import (
	"encoding/json"
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/auth"
	"github.com/kakkky/app/domain/errors"
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

// @Summary     ユーザーのログイン
// @Description メールアドレス・パスワードで認証し、署名されたトークンを返す
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       request body     LoginRequest                             true "認証に必要な情報"
// @Success     200     {object} presenter.SuccessResponse[LoginResponse] "署名されたトークンを含む情報"
// @Failure     400     {object} presenter.FailureResponse                "不正なリクエスト"
// @Failure     401     {object} presenter.FailureResponse                "パスワードが不一致"
// @Failure     500     {object} presenter.FailureResponse                "内部サーバーエラー"
// @Router      /login [post]
func (lh *LoginHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var params LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidator().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	input := auth.LoginUsecaseInputDTO{
		Email:    params.Email,
		Password: params.Password,
	}
	output, err := lh.loginUsecase.Run(r.Context(), input)
	if (err != nil) && errors.Is(err, errors.ErrPasswordMismatch) || errors.IsDomainErr(err) {
		presenter.RespondUnAuthorized(rw, "email or password invalid")
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := LoginResponse{
		JwtToken: output.JwtToken,
	}
	presenter.RespondOK(rw, resp)
}
