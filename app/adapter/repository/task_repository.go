package repository

import (
	"context"
	"database/sql"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/task"
)

type taskRepository struct {
	querier Querier
}

func NewTaskRepository(querier Querier) task.TaskRepository {
	return &taskRepository{
		querier: querier,
	}
}

func (tr *taskRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
	t, err := tr.querier.FindTaskById(ctx, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, errors.ErrNotFoundTask
	}
	if err != nil {
		return nil, err
	}
	task := task.ReconstructTask(
		t.ID,
		t.UserID,
		t.Content,
		int(t.State),
	)
	return task, nil
}

func (tr *taskRepository) Save(ctx context.Context, task *task.Task) error {
	arg := InsertTaskParams{
		ID:      task.GetID(),
		UserID:  task.GetUserId(),
		Content: task.GetContent().Value(),
		State:   int32(task.GetState().IntValue()),
	}
	if err := tr.querier.InsertTask(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Update(ctx context.Context, task *task.Task) error {
	arg := UpdateTaskParams{
		ID:    task.GetID(),
		State: int32(task.GetState().IntValue()),
	}
	if err := tr.querier.UpdateTask(ctx, arg); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Delete(ctx context.Context, task *task.Task) error {
	if err := tr.querier.DeleteTask(ctx, task.GetID()); err != nil {
		return err
	}
	return nil
}
