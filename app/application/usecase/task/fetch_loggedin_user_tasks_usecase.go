package task

import "context"

type FetchLoggedInUserTasksUsecase struct {
	taskQueryService TaskQueryService
}

func NewFetchLoggedInUserTasksUsecase(taskQueryService TaskQueryService) *FetchLoggedInUserTasksUsecase {
	return &FetchLoggedInUserTasksUsecase{
		taskQueryService: taskQueryService,
	}
}

func (fltu *FetchLoggedInUserTasksUsecase) Run(ctx context.Context, input FetchLoggedInUserTasksUsecaseInputDTO) (
	[]*FetchLoggedInUserTasksUsecaseOutputDTO, error,
) {
	dtos, err := fltu.taskQueryService.FetchTaskByUserId(ctx, input.UserId)
	if err != nil {
		return nil, err
	}
	// クエリサービスから得たdtoを詰め替える
	outputs := make([]*FetchLoggedInUserTasksUsecaseOutputDTO, 0, len(dtos))
	for _, dto := range dtos {
		outputs = append(outputs, &FetchLoggedInUserTasksUsecaseOutputDTO{
			ID:      dto.ID,
			Content: dto.Content,
			State:   dto.State,
		})
	}
	return outputs, nil
}
