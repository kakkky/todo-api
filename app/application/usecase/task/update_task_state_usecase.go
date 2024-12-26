package task

import (
	"context"

	"github.com/kakkky/app/domain/task"
)

type UpdateTaskStateUsecase struct {
	taskRepository task.TaskRepository
}

func NewUpdateTaskStateUsecase(taskRepository task.TaskRepository) *UpdateTaskStateUsecase {
	return &UpdateTaskStateUsecase{
		taskRepository: taskRepository,
	}
}

func (utu *UpdateTaskStateUsecase) Run(ctx context.Context, input UpdateTaskStateUsecaseInputDTO) (
	*UpdateTaskStateUsecaseOutputDTO, error,
) {
	t, err := utu.taskRepository.FindById(ctx, input.ID)
	if err != nil || t == nil {
		return nil, err
	}
	updatedTask, err := t.UpdateState(input.State)
	if err != nil {
		return nil, err
	}
	if err := utu.taskRepository.Update(ctx, updatedTask); err != nil {
		return nil, err
	}
	return &UpdateTaskStateUsecaseOutputDTO{
		ID:      updatedTask.GetID(),
		UserId:  updatedTask.GetUserId(),
		Content: updatedTask.GetContent().Value(),
		State:   updatedTask.GetState().Value(),
	}, nil
}
