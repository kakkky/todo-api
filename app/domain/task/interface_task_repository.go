package task

import "context"

//go:generate mockgen -package=task -source=./interface_task_repository.go -destination=./mock_task_repository.go
type TaskRepository interface {
	FindById(ctx context.Context, id string) (*Task, error)
	Save(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, task *Task) error
}
