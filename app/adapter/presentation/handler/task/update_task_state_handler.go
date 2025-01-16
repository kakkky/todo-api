package task

import (
	"encoding/json"
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/pkg/validation"
)

type UpdateTaskStateHandler struct {
	updateTaskStateUsecase *task.UpdateTaskStateUsecase
}

func NewUpdateTaskStateHandler(updateTaskStateUsecase *task.UpdateTaskStateUsecase) *UpdateTaskStateHandler {
	return &UpdateTaskStateHandler{
		updateTaskStateUsecase: updateTaskStateUsecase,
	}
}

// @Summary     タスク状態を更新する
// @Description タスクの状態(todo/doing/done)を 指定して更新する
// @Tags        Task
// @Produce     json
// @Param       request body UpdateTaskStateRequest true "タスク更新のための情報"
// @Security    BearerAuth
// @Success     201 {object} presenter.SuccessResponse[UpdateTaskStateResponse] "更新したタスクの情報"
// @Failure     400 {object} presenter.FailureResponse                          "不正なリクエスト"
// @Failure     403 {object} presenter.FailureResponse                          "権限エラー"
// @Failure     500 {object} presenter.FailureResponse                          "内部サーバーエラー"
// @Router      /tasks/{id} [patch]
func (utsh *UpdateTaskStateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// jsonをデコード
	var params UpdateTaskStateRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err := validation.NewValidation().Struct(&params); err != nil {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	// idをパスパラメータから取得
	id := r.PathValue("id")
	// contextからuserIdを取得
	userId := r.Context().Value(middleware.UserIDKey{}).(string)
	// inputDTOに詰め替える
	input := task.UpdateTaskStateUsecaseInputDTO{
		ID:     id,
		UserId: userId,
		State:  params.State,
	}
	output, err := utsh.updateTaskStateUsecase.Run(r.Context(), input)
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
	resp := UpdateTaskStateResponse{
		ID:      output.ID,
		UserId:  output.UserId,
		Content: output.Content,
		State:   output.State,
	}
	presenter.RespondOK(rw, resp)
}
