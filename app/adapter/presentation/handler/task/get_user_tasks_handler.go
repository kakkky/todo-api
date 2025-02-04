package task

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/middleware"
	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
)

type GetUserTasksHandler struct {
	fetchUserTasksUsecase *task.FetchUserTasksUsecase
}

func NewGetUserTasksHandler(fetchUserTasksUsecase *task.FetchUserTasksUsecase) *GetUserTasksHandler {
	return &GetUserTasksHandler{
		fetchUserTasksUsecase: fetchUserTasksUsecase,
	}
}

// @Summary ユーザーが持つ全てのタスクを表示する
// @Description　ログインしているユーザーのタスクを全て表示する
// @Tags     Task
// @Produce  json
// @Security BearerAuth
// @Success  200 {object} presenter.SuccessResponse[[]GetTaskResponse] "タスクの情報"
// @Failure  400 {object} presenter.FailureResponse                    "不正なリクエスト"
// @Failure  500 {object} presenter.FailureResponse                    "内部サーバーエラー"
// @Router   /user/tasks [get]
func (guth *GetUserTasksHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	input := task.FetchUserTasksUsecaseInputDTO{
		UserId: userID,
	}
	outputs, err := guth.fetchUserTasksUsecase.Run(r.Context(), input)
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := make([]GetUserTaskResponse, 0, len(outputs))
	for _, output := range outputs {
		resp = append(resp, GetUserTaskResponse{
			ID:      output.ID,
			Content: output.Content,
			State:   output.State,
		})
	}
	presenter.RespondOK(rw, resp)
}
