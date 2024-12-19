package user

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/user"
	"github.com/kakkky/app/domain/errors"
)

type DeleteUserHandler struct {
	unregisterUsecase *user.UnregisterUsecase
}

func NewDeleteUserHandler(unregisterUsecase *user.UnregisterUsecase) *DeleteUserHandler {
	return &DeleteUserHandler{
		unregisterUsecase: unregisterUsecase,
	}
}

// @Summary     ユーザーの退会
// @Description ユーザーを退会させ、ユーザー情報を削除する
// @Tags        User
// @Produce     json
// @Param       id path string true "削除するユーザーのid"
// @Success     204
// @Failure     400 {object} presenter.FailureResponse "不正なリクエスト"
// @Failure     500 {object} presenter.FailureResponse "内部サーバーエラー"
// @Router      /user/{id} [delete]
func (duh *DeleteUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.UserIDKey{}).(string)
	if id == "" {
		presenter.RespondBadRequest(w, "include the user ID as a path parameter in the URL")
		return
	}
	input := user.UnregisterUsecaseInputDTO{
		ID: id,
	}
	ctx := r.Context()
	err := duh.unregisterUsecase.Run(ctx, input)
	if (err != nil) && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(w, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(w, err.Error())
		return
	}
	presenter.RespondNoContent(w)
}
