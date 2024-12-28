package task

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
)

type DeleteTaskHandler struct {
	deleteTaskUsease *task.DeleteTaskUsecase
}

func NewDeleteTaskHandler(deleteTaskUsease *task.DeleteTaskUsecase) *DeleteTaskHandler {
	return &DeleteTaskHandler{
		deleteTaskUsease: deleteTaskUsease,
	}
}

// @Summary     タスクを削除する
// @Description 指定したidのタスクを削除する
// @Tags        Task
// @Produce     json
// @Security    BearerAuth
// @Success     204
// @Failure     400 {object} presenter.FailureResponse "不正なリクエスト"
// @Failure     403 {object} presenter.FailureResponse "権限エラー"
// @Failure     500 {object} presenter.FailureResponse "内部サーバーエラー"
// @Router      /task/{id} [delete]
func (dth *DeleteTaskHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// パスパラメータから取得
	id := r.PathValue("id")
	// contextからuserIdを取得
	userId := r.Context().Value(middleware.UserIDKey{}).(string)
	// inputDTOに詰め替える
	input := task.DeleteTaskUsecaseInputDTO{
		ID:     id,
		UserId: userId,
	}
	err := dth.deleteTaskUsease.Run(r.Context(), input)
	// タスクを削除する権限がない（ログインしているユーザーのタスクでない）場合
	if err != nil && errors.Is(err, errors.ErrForbiddenTaskOperation) {
		presenter.RespondForbidden(rw, err.Error())
		return
	}
	// ドメインエラー
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	// その他エラー
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	presenter.RespondNoContent(rw)
}
