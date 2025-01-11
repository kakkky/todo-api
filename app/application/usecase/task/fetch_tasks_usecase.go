package task

import "context"

type FetchTasksUsease struct {
	taskQueryService TaskQueryService
}

func NewFetchTasksUsease(taskQueryService TaskQueryService) *FetchTasksUsease {
	return &FetchTasksUsease{
		taskQueryService: taskQueryService,
	}
}

func (ftu *FetchTasksUsease) Run(ctx context.Context) (
	[]*FetchTaskUsecaseOutputDTO, error,
) {
	dtos, err := ftu.taskQueryService.FetchAllTasks(ctx)
	if err != nil {
		return nil, err
	}
	// クエリサービスから得たdtoを詰め替える
	outputs := make([]*FetchTaskUsecaseOutputDTO, 0, len(dtos))
	for _, dto := range dtos {
		outputs = append(outputs, &FetchTaskUsecaseOutputDTO{
			ID:       dto.ID,
			UserId:   dto.UserId,
			UserName: dto.UserName,
			Content:  dto.Content,
			State:    dto.State,
		})
	}
	return outputs, nil
}
