package task

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
)

type GetTasksHandler struct {
	fetchTasksUsecase *task.FetchTasksUsease
}

func NewGetTasksHandler(fetchTasksUsecase *task.FetchTasksUsease) *GetTasksHandler {
	return &GetTasksHandler{
		fetchTasksUsecase: fetchTasksUsecase,
	}
}

func (gth *GetTasksHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	outputs, err := gth.fetchTasksUsecase.Run(r.Context())
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := make([]GetTaskResponse, 0, len(outputs))
	for _, output := range outputs {
		resp = append(resp, GetTaskResponse{
			ID:       output.ID,
			UserId:   output.UserId,
			UserName: output.UserName,
			Content:  output.Content,
			State:    output.State,
		})
	}
	presenter.RespondOK(rw, resp)
}
