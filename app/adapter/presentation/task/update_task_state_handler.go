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
	// contextからuserIdを取得
	userId := r.Context().Value(middleware.UserIDKey{}).(string)
	// inputDTOに詰め替える
	input := task.UpdateTaskStateUsecaseInputDTO{
		ID:     params.ID,
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
