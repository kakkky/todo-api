package user

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/errors"
)

type GetUsersHandler struct {
	fetchAllUserUsecase *user.FetchUsersUsecase
}

func NewGetUsersHandler(fetchAllUserUsecase *user.FetchUsersUsecase) *GetUsersHandler {
	return &GetUsersHandler{
		fetchAllUserUsecase: fetchAllUserUsecase,
	}
}

// @Summary     全ユーザーを取得する
// @Description 全てのユーザーのID・名前をリストで取得する
// @Tags        User
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} presenter.SuccessResponse[[]GetUsersResponse] "登録されたユーザーの情報"
// @Failure     400 {object} presenter.FailureResponse                     "不正なリクエスト"
// @Failure     500 {object} presenter.FailureResponse                     "内部サーバーエラー"
// @Router      /users [get]
func (guh *GetUsersHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	outputs, err := guh.fetchAllUserUsecase.Run(ctx)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := make([]GetUserResponse, 0, len(outputs))
	for _, output := range outputs {
		resp = append(resp, GetUserResponse{
			ID:   output.ID,
			Name: output.Name,
		})
	}
	presenter.RespondOK(rw, resp)

}
