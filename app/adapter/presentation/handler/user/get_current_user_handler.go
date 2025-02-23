package user

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/errors"
)

type GetCurrentUserHandler struct {
	fetchUserUsecase *user.FetchUserUsecase
}

func NewGetCurrentUserHandler(fetchUserUsecase *user.FetchUserUsecase) *GetCurrentUserHandler {
	return &GetCurrentUserHandler{
		fetchUserUsecase: fetchUserUsecase,
	}
}

// @Summary     ログインしているユーザーを取得する
// @Description トークンを元に、ログインしているユーザー情報（id,name）を返す
// @Tags        User
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} presenter.SuccessResponse[GetUserResponse] "ユーザーの情報"
// @Failure     400 {object} presenter.FailureResponse                     "不正なリクエスト"
// @Failure     500 {object} presenter.FailureResponse                     "内部サーバーエラー"
// @Router      /users/me [get]
func (gcuh *GetCurrentUserHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// リクエストスコープからユーザーIDを取得
	userID := middleware.GetUserID(ctx)
	input := user.FetchUserUsecaseInputDTO{ID: userID}
	output, err := gcuh.fetchUserUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := GetUserResponse{ID: output.ID, Name: output.Name}
	presenter.RespondOK(rw, resp)

}
