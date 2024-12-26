package task

import "context"

type FetchTaskUsease struct {
	taskQueryService TaskQueryService
}

func NewFetchTaskUsease(taskQueryService TaskQueryService) *FetchTaskUsease {
	return &FetchTaskUsease{
		taskQueryService: taskQueryService,
	}
}

func (ftu *FetchTaskUsease) Run(ctx context.Context, input FetchTaskUsecaseInputDTO) (
	*FetchTaskUsecaseOutputDTO, error,
) {
	dto, err := ftu.taskQueryService.FetchTaskById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	// クエリサービスから得たdtoを詰め替える
	return &FetchTaskUsecaseOutputDTO{
		ID:       dto.ID,
		UserId:   dto.UserId,
		UserName: dto.UserName,
		Content:  dto.Content,
		State:    dto.State,
	}, nil
}
