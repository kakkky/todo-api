package task

import "context"

type FetchUserTasksUsecase struct {
	taskQueryService TaskQueryService
}

func NewFetchUserTasksUsecase(taskQueryService TaskQueryService) *FetchUserTasksUsecase {
	return &FetchUserTasksUsecase{
		taskQueryService: taskQueryService,
	}
}

func (fltu *FetchUserTasksUsecase) Run(ctx context.Context, input FetchUserTasksUsecaseInputDTO) (
	[]*FetchUserTasksUsecaseOutputDTO, error,
) {
	dtos, err := fltu.taskQueryService.FetchUserTasks(ctx, input.UserId)
	if err != nil {
		return nil, err
	}
	// クエリサービスから得たdtoを詰め替える
	outputs := make([]*FetchUserTasksUsecaseOutputDTO, 0, len(dtos))
	for _, dto := range dtos {
		outputs = append(outputs, &FetchUserTasksUsecaseOutputDTO{
			ID:      dto.ID,
			Content: dto.Content,
			State:   dto.State,
		})
	}
	return outputs, nil
}
