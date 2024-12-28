package task

import (
	"net/http"

	"github.com/kakkky/app/adapter/presentation/presenter"
	"github.com/kakkky/app/application/usecase/task"
	"github.com/kakkky/app/domain/errors"
)

type GetTaskHandler struct {
	fetchTaskUsecase *task.FetchTaskUsease
}

func NewGetTaskHandler(fetchTaskUsecase *task.FetchTaskUsease) *GetTaskHandler {
	return &GetTaskHandler{
		fetchTaskUsecase: fetchTaskUsecase,
	}
}

func (gth *GetTaskHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// パスパラメータから取得
	id := r.PathValue("id")
	// inputDTOに詰め替える
	input := task.FetchTaskUsecaseInputDTO{
		ID: id,
	}
	output, err := gth.fetchTaskUsecase.Run(r.Context(), input)
	if err != nil && errors.IsDomainErr(err) {
		presenter.RespondBadRequest(rw, err.Error())
		return
	}
	if err != nil {
		presenter.RespondInternalServerError(rw, err.Error())
		return
	}
	resp := GetTaskResponse{
		ID:       output.ID,
		UserId:   output.UserId,
		UserName: output.UserName,
		Content:  output.Content,
		State:    output.State,
	}
	presenter.RespondOK(rw, resp)
}
