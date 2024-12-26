package repository

import (
	"context"
	"database/sql"

	"github.com/kakkky/app/domain/errors"
	"github.com/kakkky/app/domain/task"
	"github.com/kakkky/app/infrastructure/db/sqlc"
)

type taskRepository struct{}

func NewTaskRepository() task.TaskRepository {
	return &taskRepository{}
}

func (tr *taskRepository) FindById(ctx context.Context, id string) (*task.Task, error) {
	queries := sqlc.GetQueries()
	t, err := queries.FindTaskById(ctx, id)
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
	queries := sqlc.GetQueries()
	params := sqlc.InsertTaskParams{
		ID:      task.GetID(),
		UserID:  task.GetUserId(),
		Content: task.GetContent().Value(),
		State:   int32(task.GetState().IntValue()),
	}
	if err := queries.InsertTask(ctx, params); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Update(ctx context.Context, task *task.Task) error {
	queries := sqlc.GetQueries()
	params := sqlc.UpdateTaskParams{
		ID:    task.GetID(),
		State: int32(task.GetState().IntValue()),
	}
	if err := queries.UpdateTask(ctx, params); err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Delete(ctx context.Context, task *task.Task) error {
	queries := sqlc.GetQueries()
	if err := queries.DeleteTask(ctx, task.GetID()); err != nil {
		return err
	}
	return nil
}
